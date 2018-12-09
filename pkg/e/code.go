package e

const (
	SUCCESS            = 200
	SUCCESS_total_task = 201
	SUCCESS_sub_task   = 202
	SUCCESS3           = 203
	SUCCESS4           = 204
	ERROR              = 500
	INVALID_PARAMS     = 400

	ERROR_NOT_EXIST_TASK = 10001

	// 保存CSV失败
	ERROR_UPLOAD_SAVE_CSV_FAIL = 30001
	// 检查CSV失败
	ERROR_UPLOAD_CHECK_CSV_FAIL = 30002
	// 校验CSV错误，CSV格式或大小有问题
	ERROR_UPLOAD_CHECK_CSV_FORMAT = 30003
)
