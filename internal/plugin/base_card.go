package plugin

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"os"

	"github.com/dronestock/drone/internal/config"
	"github.com/goexl/exc"
)

func (b *Base) WriteCard(url string, _card any) (err error) {
	if nil == _card {
		return
	}

	__card := new(config.CardOutput)
	__card.Schema = url
	if __card.Data, err = json.Marshal(_card); nil != err {
		return
	}

	if data, je := json.Marshal(__card); nil == je {
		switch b.Card.Path {
		case "/dev/stdout":
			err = b.writeCardTo(os.Stdout, data)
		case "/dev/stderr":
			err = b.writeCardTo(os.Stderr, data)
		case "":
			err = exc.NewMessage(`卡片写入路径为空`)
		default:
			err = os.WriteFile(b.Card.Path, data, 0600)
		}
	} else {
		err = je
	}

	return
}

func (b *Base) writeCardTo(out io.Writer, data []byte) (err error) {
	encoded := base64.StdEncoding.EncodeToString(data)
	if _, err = io.WriteString(out, "\u001B]1338;"); nil != err {
		return
	}
	if _, err = io.WriteString(out, encoded); nil != err {
		return
	}
	if _, err = io.WriteString(out, "\u001B]0m"); nil != err {
		return
	}
	_, err = io.WriteString(out, "\n")

	return
}