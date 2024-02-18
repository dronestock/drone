package plugin

import (
	"github.com/dronestock/drone/internal/config"
	"github.com/go-resty/resty/v2"
	"github.com/goexl/http"
)

func (b *Base) Http() *http.Client {
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
	client := http.New()
	if nil == b.Proxies {
		b.Proxies = make([]*config.Proxy, 0)
	}
	if nil != b.Proxy {
		b.Proxies = append(b.Proxies, b.Proxy)
	}

	for _, _proxy := range b.Proxies {
		builder := client.Proxy()
		builder.Host(_proxy.Host)
		builder.Scheme(_proxy.Scheme)
		builder.Port(_proxy.Port)
		builder.Target(_proxy.Target)
		builder.Exclude(_proxy.Exclude)
		client = builder.Build()
	}
	b.http = client.Build()
}
