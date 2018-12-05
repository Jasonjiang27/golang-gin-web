package v1

import (
	"log"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"  //golang数据校验的一个包
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

//提交csv任务
func TaskSubmit(c *gin.Context) {

<<<<<<< HEAD
=======
	//上传csv文件
	//file, _ := c.FormFile("file")
	//log.Println(file.Filename)

	taskType := c.Query("task_type")
	fileName := c.Query("file_name")
	projectName := c.Query("project_name")
	columnNumber := com.StrTo(c.Query("column_number")).MustInt()
	isAppend := c.Query("is_append")
	numberLables := com.StrTo(c.Query("number_labels")).MustInt()
	lineNumbers := com.StrTo(c.Query("line_numbers")).MustInt()

	valid := validation.Validation{}
	valid.Required(taskType, "task_type").Message("任务类型不呢个为空")
	valid.Required(fileName, "file_name").Message("上传的文件不能为空")
	valid.Required(projectName, "project_name").Message("分类树名不能为空")
	valid.Required(columnNumber, "column_number").Message("处理的列数不能为空")
	valid.Required(isAppend, "is_append").Message("添加不能为空，只能是或否")
	valid.Range(numberLables, 0, 1, "number_labels").Message("数字标签只能是0或1") //0 代表单个标签多行拆分,1 代表多个标签多行拆分
	valid.Min(lineNumbers, 0, "line_numbers").Message("拆分任务数不能为空")

	//taskId := com.StrTo(c.Query("task_id")).MustInt()
	userId := com.StrTo(c.Query("user_id")).MustInt()
	fileLocation := c.Query("file_location")
	limit := com.StrTo(c.Query("limit")).MustInt()
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	subTaskNumbers := com.StrTo(c.Query("sub_task_numbers")).MustInt()
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		
		data := make(map[string]interface{})
		//data["task_id"] = taskId
		data["user_id"] = userId
		data["task_type"] = taskType
		data["file_name"] = fileName
		data["file_location"] = fileLocation
		data["task_project_name"] = projectName

		data["task_column_number"] = columnNumber
		//data["task_status"] = c.Query("task_status")
		data["limit"] = limit

		data["task_type"] = taskType
		data["file_name"] = fileName
		data["start_time"] = startTime
		data["end_time"] = endTime
		data["sub_task_numbers"] = subTaskNumbers
		
		models.TaskSubmit(data)

		code = e.SUCCESS
		
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})

>>>>>>> 71f05493a1d898ff4204b51cfd5154454e5721d3
}

//提交mongo任务
func TaskCommonSubmit(c *gin.Context) {

<<<<<<< HEAD
=======
	//taskType := "common"
	//source := getDataSource(c *gin.Context)

	//brand, series := models.GetBrands()
	//limit := 100


	//startTime := time.Now().Unix()
	
    //执行跑批测试
	//endTime := time.Now().Unix()
	taskType := c.Query("task_type")
	
	projectName := c.Query("project_name")
	columnNumber := com.StrTo(c.Query("column_number")).MustInt()



	userId := com.StrTo(c.Query("user_id")).MustInt()
	fileLocation := c.Query("file_location")
	limit := com.StrTo(c.Query("limit")).MustInt()
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	subTaskNumbers := com.StrTo(c.Query("sub_task_numbers")).MustInt()


	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	data["user_id"] = userId
	data["task_type"] = taskType

	data["file_location"] = fileLocation
	data["task_project_name"] = projectName

	data["task_column_number"] = columnNumber
	//data["task_status"] = c.Query("task_status")
	data["limit"] = limit

	data["start_time"] = startTime
	data["end_time"] = endTime
	data["sub_task_numbers"] = subTaskNumbers
		
	//data["task_status"] = c.Query("task_status")
	data["limit"] = limit



	data["start_time"] = startTime
	data["end_time"] = endTime
	data["sub_task_numbers"] = subTaskNumbers
	
	models.TaskCommonSubmit(data)
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}

//执行跑批测试
func TaskTest(){

>>>>>>> 71f05493a1d898ff4204b51cfd5154454e5721d3
}

//获取任务列表
func GetTasks(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}  //数据校验功能

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
<<<<<<< HEAD

=======
	task_id := com.StrTo(c.Param("task_id")).MustInt()

	valid := validation.Validation{}
	valid.Min(task_id, 1, "task_id").Message("任务id必须大于0")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		
		models.DeleteTask(task_id)
		code = e.SUCCESS
		data["删除的任务"] = task_id
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,

	})
>>>>>>> 71f05493a1d898ff4204b51cfd5154454e5721d3
}
