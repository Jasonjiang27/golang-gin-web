package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Unknwon/com"

	"golang-gin-web/pkg/e"
	"golang-gin-web/models"
	"golang-gin-web/pkg/util"
	"golang-gin-web/pkg/setting"
)

//获取数据来源
func GetDataSource(c *gin.Context){

}

//获取品牌
func GetBrands(c *gin.Context) {

}

//跑批结果文件下载
func DownFile(c *gin.Context) {

}

//查看任务进度
func TaskProcess(c *gin.Context){

}

//提交csv任务
func TaskSubmit(c *gin.Context) {

}

//提交mongo任务
func TaskCommonSubmit(c *gin.Context) {

}

//获取任务列表
func GetTask(c *gin.Context) {
	task_id := c.Query("taskId")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if task_id > 0 {
		maps[task_id] = task_id
	}
	if user_id > 0 {
		maps[user_id] = user_id
	}
	maps[task_status] = task_status
	
}

//跑批任务删除
func DeleteTask(c *gin.Context) {

}