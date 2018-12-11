package v1

import (
	"fmt"
	//"time"
	//"math"
	"log"
	"net/http"

	//"strings"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation" //golang数据校验的一个包
	"github.com/gin-gonic/gin"

	"golang-gin-web/models"
	"golang-gin-web/pkg/e"
	//"golang-gin-web/pkg/setting"
	"golang-gin-web/pkg/util"
	//"golang-gin-web/pkg/mongodb"
)

//获取数据来源
func GetDataSource(c *gin.Context) {

}

//获取品牌
func GetBrands(c *gin.Context) {

}

//查看任务进度
func TaskProcess(c *gin.Context) {
	task_id := c.Param("task_id")
	//task_uid := c.Query("task_uid")
	log.Println(task_id)
	code := e.INVALID_PARAMS
	data1 := make(map[string]interface{})
	data2 := make(map[string]interface{})

	data1["task_id"] = task_id
	data2["task_id"] = task_id
	data2["status"] = "success"

	sub_tasks_total := models.SubTaskCount(data1)          //子任务总数
	sub_tasks_done_total := models.SubTaskDoneCount(data2) //子任务完成数
	log.Printf("总任务数：%d,子任务完成数：%d", sub_tasks_total, sub_tasks_done_total)
	task_process := float64(sub_tasks_done_total) / float64(sub_tasks_total) //进度

	s := fmt.Sprintf("%.2f", task_process) //保留2位小数

	code = e.SUCCESS
	msg := make(map[string]interface{})
	msg["task_id"] = task_id
	msg["process"] = s

	c.JSON(http.StatusOK, gin.H{
		"errorcode": code,
		"msg":  msg,
	})

}

//提交csv任务
func TaskSubmit(c *gin.Context) {

	task_type := "csv"
	file_name := c.Query("file_name")
	task_project_name := c.Query("task_project_name")
	task_column_number := com.StrTo(c.Query("task_column_number")).MustInt()
	is_append := c.Query("is_append")
	number_lables := com.StrTo(c.Query("number_labels")).MustInt()
	line_numbers := com.StrTo(c.Query("line_numbers")).MustInt()

	valid := validation.Validation{}
	valid.Required(task_type, "task_type").Message("任务类型不能为空")
	valid.Required(file_name, "file_name").Message("上传的文件不能为空")
	valid.Required(task_project_name, "task_project_name").Message("分类树名不能为空")
	valid.Required(task_column_number, "task_column_number").Message("处理的列数不能为空")
	valid.Required(is_append, "is_append").Message("添加不能为空，只能是或否")
	valid.Range(number_lables, 0, 1, "number_labels").Message("数字标签只能是0或1") //0 代表单个标签多行拆分,1 代表多个标签多行拆分
	valid.Min(line_numbers, 0, "line_numbers").Message("拆分任务数不能为空")

	//taskId := com.StrTo(c.Query("task_id")).MustInt()
	user_id := com.StrTo(c.Query("user_id")).MustInt()
	//limit := com.StrTo(c.Query("limit")).MustInt()
	task_status := c.Query("task_status")
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {

		//数据插入总任务表
		data_task := make(map[string]interface{})
		data_task["file_name"] = file_name
		data_task["user_id"] = user_id
		data_task["task_type"] = task_type
		data_task["task_project_name"] = task_project_name
		data_task["task_column_number"] = task_column_number
		data_task["task_status"] = task_status
		//data_task["limit"] = limit

		//data_sub_task := make(map[string]interface{})

		// data_sub_task["task_project_name"] = task_project_name

		// data_sub_task["task_type"] = task_type

		//根据参数启动csv文件操作
		CsvHandle(data_task)
		//code2 = e.SUCCESS_sub_task
		code = e.SUCCESS_ADD

	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": code,
		"msg":   e.GetMsg(code),
		//"data":  make(map[string]interface{}),
	})

}

//提交mongo任务
func TaskCommonSubmit(c *gin.Context) {

	task_type := "common"
	data_source := c.Query("data_source")
	brand := c.Query("brand")
	series := c.Query("series")
	limit := com.StrTo(c.Query("limit")).MustInt()
	task_project_name := c.Query("task_project_name")

	line_numbers := com.StrTo(c.Query("line_numbers")).MustInt()
	time_from := c.Query("time_from")
	time_to := c.Query("time_to")

	valid := validation.Validation{}
	valid.Required(task_type, "task_type").Message("任务类型不能为空")
	valid.Required(data_source, "data_source").Message("数据来源不能为空")
	valid.Required(brand, "brand").Message("车品牌不能为空")
	valid.Required(series, "series").Message("车系不能为空")
	valid.Required(task_project_name, "task_project_name").Message("分类树名不能为空")

	valid.Min(line_numbers, 0, "line_numbers").Message("拆分任务数不能为空,0代表不拆分")
	valid.Required(time_to, "time_to").Message("任务筛选结束时间不能为空")
	valid.Required(time_from, "time_from").Message("任务筛选起始时间不能为空")

	user_id := com.StrTo(c.Query("user_id")).MustInt()
	//file_location := c.Query("file_location")

	start_time := c.Query("start_time")
	end_time := c.Query("end_time")

	task_status := c.Query("task_status")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		data_task := make(map[string]interface{})
		data_task["user_id"] = user_id
		data_task["task_type"] = "common"
		data_task["data_source"] = data_source
		//data["file_location"] = file_location
		data_task["task_project_name"] = task_project_name

		//data["task_status"] = c.Query("task_status")
		data_task["limit"] = limit
		data_task["task_status"] = task_status
		data_task["start_time"] = start_time
		data_task["end_time"] = end_time
		data_task["task_id"] = util.EncodeMD5(data_source + brand + series + task_project_name)

		//数据插入子任务表
		data_mongo := make(map[string]interface{})
		data_mongo["k_source"] = data_source
		data_mongo["k_c_brand"] = brand
		data_mongo["k_c_set"] = series

		data_task["sub_task_numbers"] = models.CountData(data_mongo)
		log.Println(data_source)
		log.Println(brand)
		log.Println(series)
		log.Println(data_task["sub_task_numbers"])
		//插入数据至总任务表
		models.TaskCommonSubmit(data_task)
		
		task_texts, err := models.FindData(data_mongo) //获取mongo表中的所有文本内容数据
		if err != nil {
			panic(err)
		} else {
			for i, task_text := range task_texts {
				data_sub_task := make(map[string]interface{})
				data_sub_task["task_id"] = data_task["task_id"] //TODO:需要关联至总任务表中的task_id
				data_sub_task["task_text"] = task_text.Content
				data_sub_task["task_project_name"] = task_project_name
				data_sub_task["task_type"] = task_type
				data_sub_task["number_id"] = i
				log.Println(data_sub_task["task_id"])
				log.Println(data_sub_task["task_text"])
				models.AddSubTask(data_sub_task)
			}
		
			//code2 = e.SUCCESS_sub_task
			code = e.SUCCESS
		}
		
		code = e.SUCCESS_ADD
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorcode": code,
		"msg":  e.GetMsg(code),
		//"data": make(map[string]interface{}),
	})
}

//获取任务列表
func GetTasks(c *gin.Context) {
	maps := make(map[string]interface{})
	
	msg := make(map[string]interface{})
	//valid := validation.Validation{} //数据校验功能

	limit := com.StrTo(c.Param("limit")).MustInt()
	user_id := com.StrTo(c.Param("user_id")).MustInt()
	task_type := c.Param("task_type")
	offset := com.StrTo(c.Param("offset")).MustInt()

	
	maps["user_id"] = user_id

	// if !valid.HasErrors() {

	// 	data["list"] = models.GetTasks(util.GetPage(c), setting.PageSize, maps)
	// 	data["total"] = models.GetTasksTotal(maps)

	// } else {
	// 	for _, err := range valid.Errors {
	// 		log.Fatal(err.Key, err.Message)

	// 	}
	// }
	msg["total"] = models.GetTasksTotal(maps)

	msg["limit"] = limit
	msg["offset"] = offset
	msg["user_id"] = user_id
	msg["status"] = e.SUCCESS
	msg["task_type"] = task_type
	msg["tasks"] = models.GetTasks(offset, limit, maps)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"errorcode":     code,
		"msg":      msg,
	})
}

//跑批任务删除
func DeleteTask(c *gin.Context) {

	task_id := c.Param("task_id")

	valid := validation.Validation{}
	valid.Required(task_id,"task_id").Message("任务不能为空")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {

		models.DeleteTask(task_id)
		code = e.SUCCESS
		data["删除的任务"] = task_id
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorcode": code,
		"msg":  e.GetMsg(code),
		//"data": data,
	})

}
