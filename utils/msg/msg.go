package msg

const (
	SUCCSE        = 200
	ERROR         = 500
	InvalidParams = 400


)

var codeMsg = map[int]string{
	SUCCSE:        "OK",
	ERROR:         "FAIL",
	InvalidParams: "请求参数错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
