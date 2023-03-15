package drone

import (
	"time"

	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

var _ = New

type bootstrap struct {
	*Base

	getter    *getter
	processor *processor
	plugin    Plugin
	name      string
	aliases   []*alias

	started time.Time
}

func New(constructor constructor) (boot *bootstrap) {
	boot = new(bootstrap)
	boot.plugin = constructor()
	boot.getter = newGetter(boot)
	boot.processor = newProcessor()

	base := new(Base)
	base.Logger = simaqian.Default()
	boot.Base = base

	return
}

func (b *bootstrap) Name(name string) *bootstrap {
	b.name = name

	return b
}

func (b *bootstrap) Alias(name string, value string) *bootstrap {
	b.aliases = append(b.aliases, newAlias(name, value))

	return b
}

func (b *bootstrap) Boot() {
	var err error
	defer b.finally(&err)

	b.started = time.Now()
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
