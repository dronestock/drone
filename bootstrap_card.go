package drone

import (
	"fmt"
	"strings"
	"time"

	"github.com/goexl/gox/field"
)

func (b *bootstrap) startCard() {
	ticker := time.NewTimer(100 * time.Millisecond)
	defer func() {
		_ = ticker.Stop()
	}()

	for range ticker.C {
		if err := b.writeCard(); nil != err {
			b.Warn("写入卡片数据出错", field.Error(err))
		}
		ticker.Reset(time.Second)
	}
}

func (b *bootstrap) writeCard() (err error) {
	scheme := b.plugin.Scheme()
	if strings.HasPrefix(scheme, github) {
		scheme = fmt.Sprintf("%s%s", ghproxy, scheme)
	}

	if _card, ce := b.plugin.Carding(); nil != ce {
		err = ce
	} else {
		err = b.Base.writeCard(scheme, _card)
	}

	return
}
