package drone

import (
	"fmt"
	"os"
	"strings"
)

func (b *bootstrap) prepare() (err error) {
	if "true" == b.getter.Get("VERBOSE") { // 执行这一步的时候，不能使用结构体里面的配置，因为还没有加载配置
		fmt.Println(strings.Repeat("-", 120))
		for _, env := range os.Environ() {
			fmt.Println(env)
		}
		fmt.Println(strings.Repeat("-", 120))
	}

	return
}
