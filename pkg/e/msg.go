package e

var MsgFlags = map[int]string {
	SUCCESS : "OK",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
}

//获取错误码
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}