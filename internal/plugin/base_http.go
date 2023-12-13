package plugin

import (
	"github.com/go-resty/resty/v2"
)

func (b *Base) Http() *resty.Client {
	if nil == b.http {
		b.setupHttp()
	}

	return b.http
}

func (b *Base) Request() *resty.Request {
	if nil == b.http {
		b.setupHttp()
	}

	return b.http.R()
}

func (b *Base) setupHttp() {
	b.http = resty.New()

	if nil != b.Proxy {
		b.http.SetProxy(b.Proxy.Addr())
	}
}
