package core

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/drone/envsubst"
	"github.com/dronestock/drone/internal"
	"github.com/goexl/env"
	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/check"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

type Getter struct {
	simaqian.Logger
	// plugin.Plugin
	// base.Base

	vm      *vm.VM
	options []expr.Option
}

func NewGetter(logger simaqian.Logger) (getter *Getter) {
	getter = new(Getter)
	getter.Logger = logger
	getter.vm = new(vm.VM)
	getter.options = []expr.Option{
		expr.AllowUndefinedVariables(),
		expr.Function(internal.FuncFile, getter.file),
		expr.Function(internal.FuncUrl, getter.url),
		expr.Function(internal.FuncHttp, getter.url),
		expr.Function(internal.FuncMatch, getter.match),
	}
	/*for _, expression := range bootstrap.Expressions() {
		getter.options = append(getter.options, expr.Function(expression.Name(), expression.Exec))
	}*/

	return
}

func (g *Getter) Get(key string) (value string) {
	defer g.recover()
	if got := g.env(key); "" != strings.TrimSpace(got) {
		value = got
	}
	if got := g.eval(value); "" != strings.TrimSpace(got) {
		value = got
	}
	if "" == value { // 如果环境变量取值没有改变，证明键没有环境变量，需要将键值赋值
		value = key
	}

	size := len(value)
	if internal.JsonObjectStart == (value)[0:1] && internal.JsonObjectEnd == (value)[size-1:size] {
		value = g.fixJsonObject(value)
	} else if internal.JsonArrayStart == (value)[0:1] && internal.JsonArrayEnd == (value)[size-1:size] {
		value = g.fixJsonArray(value)
	} else {
		value = g.expr(value)
	}

	// 如果没有一点变化，证明没有任何配置，返回空值
	if value == key {
		value = ""
	}

	return
}

func (g *Getter) expr(from string) (to string) {
	fields := gox.Fields[any]{
		field.New("Expression", from),
	}
	if program, ce := expr.Compile(from, g.options...); nil != ce {
		to = from
		g.Debug("表达式编译出错", fields.Add(field.Error(ce))...)
	} else if result, re := g.vm.Run(program, nil); nil != re {
		to = from
		g.Debug("表达式运算出错", fields.Add(field.Error(re))...)
	} else {
		to = gox.ToString(result)
		g.Debug("表达式运算成功", fields.Add(field.New("result", to))...)
	}

	return
}

func (g *Getter) fixJsonObject(from string) (to string) {
	object := make(map[string]any)
	if ue := json.Unmarshal([]byte(from), &object); nil != ue {
		to = from
	} else {
		g.fixObjectExpr(object)
	}

	if from == to {
		// 不需要进行转换
	} else if bytes, me := json.Marshal(object); nil != me {
		to = from
	} else {
		to = string(bytes)
	}

	return
}

func (g *Getter) fixJsonArray(from string) (to string) {
	array := make([]any, 0)
	if ue := json.Unmarshal([]byte(from), &array); nil != ue {
		to = from
	} else {
		g.fixArrayExpr(&array)
	}

	if from == to {
		// 不需要进行转换
	} else if bytes, me := json.Marshal(array); nil != me {
		to = from
	} else {
		to = string(bytes)
	}

	return
}

func (g *Getter) fixArrayExpr(array *[]any) {
	for index, value := range *array {
		switch vt := value.(type) {
		case string:
			(*array)[index] = g.expr(vt)
		case []any:
			g.fixArrayExpr(&vt)
		case map[string]any:
			g.fixObjectExpr(vt)
		}
	}
}

func (g *Getter) fixObjectExpr(object map[string]any) {
	for key, value := range object {
		switch vt := value.(type) {
		case string:
			object[key] = g.expr(vt)
		case []any:
			g.fixArrayExpr(&vt)
		case map[string]any:
			g.fixObjectExpr(vt)
		}
	}
}

func (g *Getter) env(key string) (value string) {
	key = strings.ToUpper(key)
	if value = os.Getenv(key); "" != value {
		return
	}
	if value = env.Get(internal.DroneEnv(key)); "" != value {
		return
	}
	if value = env.Get(key); "" != value {
		return
	}

	return
}

func (g *Getter) eval(from string) (to string) {
	to = from
	if !strings.Contains(to, internal.Dollar) {
		return
	}

	count := 0
	for {
		if value, ee := envsubst.Eval(to, g.env); nil == ee {
			to = value
		}

		if count >= 2 || !strings.Contains(to, internal.Dollar) {
			break
		}
		if strings.Contains(to, internal.Dollar) {
			count++
		}
	}

	return
}

func (g *Getter) isHttp(url string) bool {
	return check.New().
		Any().
		String(url).
		Items(internal.PrefixHttpProtocol, internal.PrefixHttpsProtocol).
		Prefix().
		Check()
}

func (g *Getter) recover() {
	if ctx := recover(); nil != ctx {
		switch value := ctx.(type) {
		case error:
			g.Warn("获取器执行出错", field.Error(value))
		}
	}
}

func (g *Getter) file(args ...any) (result any, err error) {
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
		g.Error("读取文件出错", fields.Add(field.Error(re))...)
	} else {
		result = string(bytes)
		g.Debug("读取文件成功", fields.Add(field.New("content", result))...)
	}

	return
}

func (g *Getter) url(args ...any) (result any, err error) {
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

	/*fields := gox.Fields[any]{
		field.New("url", url),
	}
	if rsp, re := g.Http().Get(url); nil != re {
		g.Error("读取端点出错", fields.Add(field.Error(re))...)
	} else if rsp.IsError() {
		httpFields := gox.Fields[any]{
			field.New("code", rsp.StatusCode()),
			field.New("body", rsp.Body()),
		}
		g.Warn("远端服务器返回错误", fields.Add(httpFields...)...)
	} else {
		result = string(rsp.Body())
		g.Debug("读取端点成功", fields.Add(field.New("content", result))...)
	}*/

	return
}

func (g *Getter) match(args ...any) (result any, err error) {
	if 2 != len(args) {
		err = exc.NewFields("参数错误", field.New("args", args), field.New("need", 2), field.New("real", 1))
	}
	if nil != err {
		return
	}

	reg := regexp.MustCompile(args[1].(string))
	result = reg.FindStringSubmatch(args[0].(string))

	return
}
