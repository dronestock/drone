package config

import (
	"fmt"
	"net/url"
)

type Proxy struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled,omitempty" yaml:"enabled" xml:"enabled" toml:"enabled"`
	// 主机
	Host string `json:"host,omitempty" yaml:"host" xml:"host" toml:"host" validate:"required"`
	// 端口
	Port int `json:"port,omitempty" yaml:"port" xml:"port" toml:"port" validate:"max=65535"`
	// 代理类型
	// nolint:lll
	Scheme string `default:"scheme,omitempty" json:"scheme" yaml:"scheme" xml:"scheme" toml:"scheme" validate:"required,oneof=socks4 socks5 http https"`
	// 目标
	Target string `json:"target,omitempty" yaml:"target" xml:"target" toml:"target"`
	// 排除
	Exclude string `json:"exclude,omitempty" yaml:"exclude" xml:"exclude" toml:"exclude"`
	// 代理认证用户名
	Username string `json:"username,omitempty" yaml:"username" xml:"username" toml:"username"`
	// 代理认证密码
	Password string `json:"password,omitempty" yaml:"password" xml:"password" toml:"password"`
}

func (p *Proxy) Addr() (addr string) {
	if "" != p.Username && "" != p.Password {
		addr = fmt.Sprintf(
			"%s://%s:%s@%s",
			p.Scheme,
			url.QueryEscape(p.Username), url.QueryEscape(p.Password),
			p.Host,
		)
	} else {
		addr = fmt.Sprintf("%s://%s", p.Scheme, p.Host)
	}

	return
}
