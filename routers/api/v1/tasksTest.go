package v1

import (
	//"time"
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
func GetDataSource(c *gin.Context){
	
	//var data []string
	//data = models.GetDataSource()
	//return data
}

//获取品牌
func GetBrands(c *gin.Context) {
	
	//name := c.Query("name")
	//series := c.Query("series")
	//var brands 
	//data["brand"] = append(data["brand"], brands)
	//return data
}

//跑批结果文件下载
func DownFile(c *gin.Context) {

}

//查看任务进度
func TaskProcess(c *gin.Context) {

}


//提交csv任务
func TaskSubmit(c *gin.Context) {

	//上传csv文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

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

	data := make(map[string]interface{})
	data["task_type"] = taskType
	data["file_name"] = fileName
	data["project_name"] = projectName
	data["column_number"] = columnNumber
	data["is_append"] = isAppend
	data["number_labels"] = numberLables
	data["line_numbers"] = lineNumbers

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})

	
}

//提交mongo任务
func TaskCommonSubmit(c *gin.Context) {

	//taskType := "common"
	//source := getDataSource(c *gin.Context)

	//brand, series := models.GetBrands()
	//limit := 100


	//startTime := time.Now().Unix()
	
    //执行跑批测试
	//endTime := time.Now().Unix()
	
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
