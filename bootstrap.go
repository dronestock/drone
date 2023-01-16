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
	if se := b.setup(); nil != se {
		err = se
		b.Error("配置插件失败", field.Error(se))
	} else if ee := b.exec(); nil != ee {
		err = ee
		b.Error("执行插件失败", field.Error(ee))
	}

	return
}
