package e

var MsgFlags = map[int]string {
	SUCCESS : "OK",
	SUCCESS_total_task : "成功添加总任务表",
	SUCCESS_sub_task : "成功添加子任务表",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",

	ERROR_NOT_EXIST_TASK : "任务不存在",

	ERROR_UPLOAD_SAVE_CSV_FAIL:    "保存CSV失败",
	ERROR_UPLOAD_CHECK_CSV_FAIL:   "检查CSV失败",
	ERROR_UPLOAD_CHECK_CSV_FORMAT: "校验CSV错误，CSV格式或大小有问题",
}

//获取错误码
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}