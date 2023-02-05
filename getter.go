package drone

import (
	"os"
	"strings"

	"github.com/goexl/env"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type getter struct {
	bootstrap *bootstrap
}

func newGetter(bootstrap *bootstrap) *getter {
	return &getter{
		bootstrap: bootstrap,
	}
}

func (g *getter) Get(key string) (value string) {
	value = g.env(key)
	if strings.HasPrefix(value, filePrefix) {
		value = g.file(strings.TrimPrefix(value, filePrefix))
	} else if strings.HasPrefix(value, urlPrefix) {
		value = g.url(strings.TrimPrefix(value, urlPrefix))
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

func (g *getter) file(name string) (value string) {
	if _, se := os.Stat(name); nil != se && os.IsNotExist(se) {
		return
	}

	fields := gox.Fields[any]{
		field.New("filename", name),
	}
	if bytes, re := os.ReadFile(name); nil != re {
		g.bootstrap.Error("读取文件出错", fields.Connect(field.Error(re))...)
	} else {
		value = string(bytes)
		g.bootstrap.Debug("读取文件成功", fields.Connect(field.New("content", value))...)
	}

	return
}

func (g *getter) url(url string) (value string) {
	if !strings.HasPrefix(url, httpProtocolPrefix) || !strings.HasPrefix(url, httpsProtocolPrefix) {
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
		value = string(rsp.Body())
		g.bootstrap.Debug("读取端点成功", fields.Connect(field.New("content", value))...)
	}

	return
}
