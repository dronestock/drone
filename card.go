package drone

type card struct {
	// 路径
	Path string `default:"${DRONE_CARD_PATH=/dev/stdout}" json:"path"`
}
