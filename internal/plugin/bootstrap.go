package plugin

import (
	"github.com/dronestock/drone/internal/core"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

type Bootstrap struct {
	*Base

	getter    *core.Getter
	processor *core.Processor
	plugin    Plugin
	name      string
	aliases   []*core.Alias
}

func New(constructor Constructor) (bootstrap *Bootstrap) {
	base := new(Base)
	base.Logger = simaqian.Default()

	bootstrap = new(Bootstrap)
	bootstrap.Base = base
	bootstrap.plugin = constructor()
	bootstrap.getter = core.NewGetter(bootstrap.Logger)
	bootstrap.processor = core.NewProcessor()

	return
}

func (b *Bootstrap) Name(name string) *Bootstrap {
	b.name = name

	return b
}

func (b *Bootstrap) Alias(name string, value string) *Bootstrap {
	b.aliases = append(b.aliases, core.NewAlias(name, value))

	return b
}

func (b *Bootstrap) Boot() {
	b.gracefully()
	var err error
	defer b.finally(&err)

	if pe := b.prepared(); nil != pe {
		err = pe
		b.Error("准备插件出错", field.Error(pe))
	} else if se := b.setup(); nil != se {
		err = se
		b.Error("配置插件出错", field.Error(se))
	} else if ee := b.run(); nil != ee {
		err = ee
		b.Error("执行插件出错", field.Error(ee))
	}
}
