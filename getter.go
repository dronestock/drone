package drone

import (
	"os"
	"strings"

	"github.com/goexl/env"
	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/check"
	"github.com/goexl/gox/field"
	"github.com/maja42/goval"
)

type getter struct {
	bootstrap *bootstrap
	functions map[string]goval.ExpressionFunction
}

func newGetter(bootstrap *bootstrap) (g *getter) {
	g = new(getter)
	g.bootstrap = bootstrap
	g.functions = map[string]goval.ExpressionFunction{
		"file": g.file,
		"url":  g.url,
		"http": g.url,
	}

	return
}

func (g *getter) Get(key string) (value string) {
	value = g.env(key)
	if "" == strings.TrimSpace(value) || !g.isExpr(value) {
		return
	}

	value = strings.ReplaceAll(value, prefixExpression, "")
	value = strings.ReplaceAll(value, prefixExp, "")
	value = strings.TrimSpace(value)
	fields := gox.Fields[any]{
		field.New("key", key),
		field.New("expression", value),
	}
	eval := goval.NewEvaluator()
	if result, ee := eval.Evaluate(value, map[string]any{}, g.functions); nil != ee {
		g.bootstrap.Debug("表达式运算出错", fields.Connect(field.Error(ee))...)
	} else {
		value = gox.ToString(result)
		g.bootstrap.Debug("表达式运算成功", fields.Connect(field.New("result", value))...)
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
		g.bootstrap.Error("读取文件出错", fields.Connect(field.Error(re))...)
	} else {
		result = string(bytes)
		g.bootstrap.Debug("读取文件成功", fields.Connect(field.New("content", result))...)
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
		g.bootstrap.Error("读取端点出错", fields.Connect(field.Error(re))...)
	} else if rsp.IsError() {
		httpFields := gox.Fields[any]{
			field.New("code", rsp.StatusCode()),
			field.New("body", rsp.Body()),
		}
		g.bootstrap.Warn("远端服务器返回错误", fields.Connects(httpFields...)...)
	} else {
		result = string(rsp.Body())
		g.bootstrap.Debug("读取端点成功", fields.Connect(field.New("content", result))...)
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

func (g *getter) isExpr(expr string) bool {
	return check.New().
		Any().
		String(expr).
		Items(prefixExp, prefixExpression, prefixExpr).
		Prefix().
		Check()
}
