package drone

import (
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

var _ = New

type bootstrap struct {
	*Base

	plugin  Plugin
	name    string
	aliases []*alias
}

func New(constructor constructor) *bootstrap {
	return &bootstrap{
		Base: &Base{
			Logger: simaqian.Default(),
		},

		plugin: constructor(),
	}
}

func (b *bootstrap) Name(name string) *bootstrap {
	b.name = name

	return b
}

func (b *bootstrap) Alias(name string, value string) *bootstrap {
	b.aliases = append(b.aliases, newAlias(name, value))

	return b
}

func (b *bootstrap) Boot() (err error) {
	defer func() {
		_ = b.commands()
	}()

	if pe := b.prepare(); nil != pe {
		err = pe
		b.Error("准备插件出错", field.Error(pe))
	} else if se := b.setup(); nil != se {
		err = se
		b.Error("配置插件出错", field.Error(se))
	} else if ee := b.exec(); nil != ee {
		err = ee
		b.Error("执行插件出错", field.Error(ee))
	}

	return
}
