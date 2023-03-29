package drone

import (
	"os"
	"regexp"

	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

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

func (g *getter) match(args ...any) (result any, err error) {
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
