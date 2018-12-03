package v1

import (
	"log"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"golang-gin-web/models"
	"golang-gin-web/pkg/e"
	"golang-gin-web/pkg/setting"
	"golang-gin-web/pkg/util"
)

//获取数据来源
func GetDataSource(c *gin.Context) {

}

//获取品牌
func GetBrands(c *gin.Context) {

}

//跑批结果文件下载
func DownFile(c *gin.Context) {

}

//查看任务进度
func TaskProcess(c *gin.Context) {

}

//新增任务
func AddTask(c *gin.Context) {

}

//提交csv任务
func TaskSubmit(c *gin.Context) {

}

//提交mongo任务
func TaskCommonSubmit(c *gin.Context) {
	taskId := com.StrTo(c.Query("task_id")).MustInt()
	UserId := com.StrTo(c.Query("user_id")).MustInt()
	Type := c.Query("type")
	State := c.DefaultQuery("state", "fail")
	TaskStatus := c.DefaultQuery("task_status", "提交中")
	TaskProjectName := c.Query("task_project_name")
	SubTaskNumbers := com.StrTo(c.Query("sub_task_numbers")).MustInt()

	valid := validation.Validation{}
	valid.Min(taskId, 1, "task_id").Message("任务id必须大于0")
	valid.Min(UserId, 1, "user_id").Message("用户id必须大于0")
	valid.Required(Type, "type").Message("数据类型不能为空")
	valid.Range(TaskStatus, "提交中", "提交成功", "task_status").Message("任务执行状态只能是提交中或任务完成")
	valid.Range(State, "成功", "失败", "state").Message("任务状态只能是成功或失败")
	valid.Required(TaskProjectName, "task_project_name").Message("任务名称不能为空")
	valid.Min(SubTaskNumbers, 1,"sub_task_numbers").Message("子任务数必须大于0")


	code := e.INVALID_PARAMS

	}


	
}

//执行跑批测试
func TaskTest(){

}

//获取任务列表
func GetTasks(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	var taskId int = -1
	if arg := c.Query("task_id"); arg != "" {
		taskId = com.StrTo(arg).MustInt()
		maps["task_id"] = taskId

		valid.Min(taskId, 1, "task_id").Message("任务id必须大于0")
	}

	var userId int = -1
	if arg := c.Query("user_id"); arg != "" {
		userId = com.StrTo(arg).MustInt()
		maps["user_id"] = userId

		valid.Min(userId, 1, "user_id").Message("用户id必须大于0")
	}

	if !valid.HasErrors() {

		data["list"] = models.GetTasks(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetTasksTotal(maps)

	} else {
		for _, err := range valid.Errors {
			log.Fatal(err.Key, err.Message)

		}
	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":     code,
		"msg":      e.GetMsg(code),
		"msg_test": "cool",
		"data":     data,
	})
}

//跑批任务删除
func DeleteTask(c *gin.Context) {
	taskId := com.StrTo(c.Param("task_id")).MustInt()

	valid := validation.Validation{}
	valid.Min(taskId, 1, "task_id").Message("任务id必须0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistTaskById(taskId) {
			models.DeleteTask(taskId)
			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Fatal(err.Key, err.Message)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data" : make(map[string]string),

		})
	}
}
