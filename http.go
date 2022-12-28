package drone

import (
	"github.com/go-resty/resty/v2"
)

func (b *Base) Http() *resty.Request {
	if nil == b.http {
		b.http = resty.New()
		b.http.SetProxy(b.Proxy.addr())
	}

	return b.http.R()
}
