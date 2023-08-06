package pkg

const (
	SUCCESS  = 200
	NOTFOUND = 404
	ERROR    = 500
)

var MsgFile = map[int]string{
	SUCCESS:  "ok",
	ERROR:    "服务错误",
	NOTFOUND: "路径错误",
}

// 接口返回结构体
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

func GetMsg(code int) string {
	msg, ok := MsgFile[code]
	if !ok {
		return MsgFile[ERROR]
	}
	return msg
}
