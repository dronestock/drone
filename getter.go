package drone

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/drone/envsubst"
	"github.com/goexl/env"
	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/check"
	"github.com/goexl/gox/field"
)

type getter struct {
	bootstrap *bootstrap
	vm        *vm.VM
	options   []expr.Option
}

func newGetter(bootstrap *bootstrap) (g *getter) {
	g = new(getter)
	g.bootstrap = bootstrap
	g.vm = new(vm.VM)
	g.options = []expr.Option{
		expr.AllowUndefinedVariables(),
		expr.Function(funcFile, g.file),
		expr.Function(funcUrl, g.url),
		expr.Function(funcHttp, g.url),
	}

	return
}

func (g *getter) Get(key string) (value string) {
	value = g.env(key)
	if "" == strings.TrimSpace(value) {
		return
	}

	value = g.eval(value)
	fields := gox.Fields[any]{
		field.New("key", key),
		field.New("value", value),
		field.New("expression", value),
	}
	if program, ce := expr.Compile(value, g.options...); nil != ce {
		g.bootstrap.Debug("表达式编译出错", fields.Add(field.Error(ce))...)
	} else if result, re := g.vm.Run(program, nil); nil != re {
		g.bootstrap.Debug("表达式运算出错", fields.Add(field.Error(re))...)
	} else {
		value = gox.ToString(result)
		g.bootstrap.Debug("表达式运算成功", fields.Add(field.New("result", value))...)
	}

	return
}

func (g *getter) env(key string) (value string) {
	defer g.json(&value)

	key = strings.ToUpper(key)
	if value = os.Getenv(key); "" != value {
		return
	}
	if value = env.Get(g.bootstrap.droneEnv(key)); "" != value {
		return
	}
	if value = env.Get(key); "" != value {
		return
	}

	return
}

func (g *getter) json(from *string) {
	if "" == strings.TrimSpace(*from) {
		// 不对空字符串做处理
	} else if "true" == *from || "false" == *from {
		// 是布尔值，满足格式要求
	} else if _, err := strconv.ParseFloat(*from, 64); nil == err {
		// 是数字，满足格式要求
	} else if jsonObjectStart == (*from)[0:1] || jsonArrayStart == (*from)[0:1] {
		// 是对象或者数据，满足格式要求
	} else {
		// 是字符串，必须加双引号
		*from = fmt.Sprintf(`"%s"`, *from)
	}
}

func (g *getter) eval(from string) (to string) {
	to = from
	if !strings.Contains(to, dollar) {
		return
	}

	count := 0
	for {
		if value, ee := envsubst.Eval(to, g.env); nil == ee {
			to = value
		}

		if count >= 2 || !strings.Contains(to, dollar) {
			break
		}
		if strings.Contains(to, dollar) {
			count++
		}
	}

	return
}

func (g *getter) file(args ...any) (result any, err error) {
	name := ""
	if 0 == len(args) {
		err = exc.NewField("必须传入参数", field.New("args", args))
	} else {
		name = gox.ToString(args[0])
	}
	if nil != err {
		return
	}

	fields := gox.Fields[any]{
		field.New("filename", name),
	}
	if bytes, re := os.ReadFile(name); nil != re {
		g.bootstrap.Error("读取文件出错", fields.Add(field.Error(re))...)
	} else {
		result = string(bytes)
		g.bootstrap.Debug("读取文件成功", fields.Add(field.New("content", result))...)
	}

	return
}

func (g *getter) url(args ...any) (result any, err error) {
	url := ""
	if 0 == len(args) {
		err = exc.NewField("必须传入参数", field.New("args", args))
	} else {
		url = gox.ToString(args[0])
		err = gox.If(g.isHttp(url), exc.NewField("必须是URL地址", field.New("url", url)))
	}
	if nil != err {
		return
	}

	fields := gox.Fields[any]{
		field.New("url", url),
	}
	if rsp, re := g.bootstrap.Http().Get(url); nil != re {
		g.bootstrap.Error("读取端点出错", fields.Add(field.Error(re))...)
	} else if rsp.IsError() {
		httpFields := gox.Fields[any]{
			field.New("code", rsp.StatusCode()),
			field.New("body", rsp.Body()),
		}
		g.bootstrap.Warn("远端服务器返回错误", fields.Add(httpFields...)...)
	} else {
		result = string(rsp.Body())
		g.bootstrap.Debug("读取端点成功", fields.Add(field.New("content", result))...)
	}

	return
}

func (g *getter) isHttp(url string) bool {
	return check.New().
		Any().
		String(url).
		Items(prefixHttpProtocol, prefixHttpsProtocol).
		Prefix().
		Check()
}
