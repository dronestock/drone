package drone

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/drone/envsubst"
	"github.com/goexl/env"
	"github.com/goexl/gox"
	"github.com/goexl/gox/check"
	"github.com/goexl/gox/field"
)

type getter struct {
	bootstrap *bootstrap
	vm        *vm.VM
	options   []expr.Option
}

func newGetter(bootstrap *bootstrap) (g *getter) {
	g = new(getter)
	g.bootstrap = bootstrap
	g.vm = new(vm.VM)
	g.options = []expr.Option{
		expr.AllowUndefinedVariables(),
		expr.Function(funcFile, g.file),
		expr.Function(funcUrl, g.url),
		expr.Function(funcHttp, g.url),
		expr.Function(funcMatch, g.match),
	}
	for _, _expression := range bootstrap.plugin.Expressions() {
		g.options = append(g.options, expr.Function(_expression.Name(), _expression.Exec))
	}

	return
}

func (g *getter) Get(key string) (value string) {
	defer g.recover()
	if got := g.env(key); "" != strings.TrimSpace(got) {
		value = got
	}
	if got := g.eval(value); "" != strings.TrimSpace(got) {
		value = got
	}

	size := len(value)
	if jsonObjectStart == (value)[0:1] && jsonObjectEnd == (value)[size-1:size] {
		value = g.fixJsonObject(value)
	} else if jsonArrayStart == (value)[0:1] && jsonArrayEnd == (value)[size-1:size] {
		value = g.fixJsonArray(value)
	} else {
		value = g.expr(value)
	}

	// 如果没有一点变化，证明没有任何配置，返回空值
	if value == key {
		value = ""
	}

	return
}

func (g *getter) expr(from string) (to string) {
	fields := gox.Fields[any]{
		field.New("expression", from),
	}
	if program, ce := expr.Compile(from, g.options...); nil != ce {
		to = from
		g.bootstrap.Debug("表达式编译出错", fields.Add(field.Error(ce))...)
	} else if result, re := g.vm.Run(program, nil); nil != re {
		to = from
		g.bootstrap.Debug("表达式运算出错", fields.Add(field.Error(re))...)
	} else {
		to = gox.ToString(result)
		g.bootstrap.Debug("表达式运算成功", fields.Add(field.New("result", to))...)
	}

	return
}

func (g *getter) fixJsonObject(from string) (to string) {
	object := make(map[string]any)
	if ue := json.Unmarshal([]byte(from), &object); nil != ue {
		to = from
	} else {
		g.fixObjectExpr(object)
	}

	if from == to {
		// 不需要进行转换
	} else if bytes, me := json.Marshal(object); nil != me {
		to = from
	} else {
		to = string(bytes)
	}

	return
}

func (g *getter) fixJsonArray(from string) (to string) {
	array := make([]any, 0)
	if ue := json.Unmarshal([]byte(from), &array); nil != ue {
		to = from
	} else {
		g.fixArrayExpr(&array)
	}

	if from == to {
		// 不需要进行转换
	} else if bytes, me := json.Marshal(array); nil != me {
		to = from
	} else {
		to = string(bytes)
	}

	return
}

func (g *getter) fixArrayExpr(array *[]any) {
	for index, value := range *array {
		switch vt := value.(type) {
		case string:
			(*array)[index] = g.expr(vt)
		case []any:
			g.fixArrayExpr(&vt)
		case map[string]any:
			g.fixObjectExpr(vt)
		}
	}
}

func (g *getter) fixObjectExpr(object map[string]any) {
	for key, value := range object {
		switch vt := value.(type) {
		case string:
			object[key] = g.expr(vt)
		case []any:
			g.fixArrayExpr(&vt)
		case map[string]any:
			g.fixObjectExpr(vt)
		}
	}
}

func (g *getter) env(key string) (value string) {
	key = strings.ToUpper(key)
	if value = os.Getenv(key); "" != value {
		return
	}
	if value = env.Get(g.bootstrap.droneEnv(key)); "" != value {
		return
	}
	if value = env.Get(key); "" != value {
		return
	}

	return
}

func (g *getter) eval(from string) (to string) {
	to = from
	if !strings.Contains(to, dollar) {
		return
	}

	count := 0
	for {
		if value, ee := envsubst.Eval(to, g.env); nil == ee {
			to = value
		}

		if count >= 2 || !strings.Contains(to, dollar) {
			break
		}
		if strings.Contains(to, dollar) {
			count++
		}
	}

	return
}

func (g *getter) isHttp(url string) bool {
	return check.New().
		Any().
		String(url).
		Items(prefixHttpProtocol, prefixHttpsProtocol).
		Prefix().
		Check()
}

func (g *getter) recover() {
	if err := recover(); nil != err {
		fmt.Println(err)
	}
}
