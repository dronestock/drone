package drone

import (
	"os"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
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

	fields := gox.Fields[any]{
		field.New("key", key),
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
