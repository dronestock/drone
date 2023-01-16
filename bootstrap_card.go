package drone

import (
	"fmt"
	"strings"
	"time"

	"github.com/goexl/gox/field"
)

func (bb *bootstrapBuilder) startCard(plugin Plugin, base *Base) {
	ticker := time.NewTimer(100 * time.Millisecond)
	defer func() {
		_ = ticker.Stop()
	}()

	for range ticker.C {
		if err := bb.writeCard(plugin, base); nil != err {
			base.Warn("写入卡片数据出错", field.Error(err))
		}
		ticker.Reset(time.Second)
	}
}

func (bb *bootstrapBuilder) writeCard(plugin Plugin, base *Base) (err error) {
	scheme := plugin.Scheme()
	if strings.HasPrefix(scheme, github) {
		scheme = fmt.Sprintf("%s%s", ghproxy, scheme)
	}

	if _card, ce := plugin.Carding(); nil != ce {
		err = ce
	} else {
		err = base.writeCard(scheme, _card)
	}

	return
}
