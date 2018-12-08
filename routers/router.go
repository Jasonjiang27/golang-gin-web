package routers

import (
	"github.com/gin-gonic/gin"

	"golang-gin-web/pkg/setting"
	"golang-gin-web/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取数据来源
		apiv1.GET("/getDataSource", v1.GetDataSource)

		//获取品牌
		apiv1.GET("/getBrands", v1.GetBrands)

		//跑批结果文件下载
		apiv1.GET("downFile/:task_id", v1.DownFile)

		//查看任务进度
		apiv1.GET("/taskProcess/:task_id", v1.TaskProcess)

		//csv文件上传
		apiv1.POST("/upload", v1.UploadFile)

		//提交csv任务
		apiv1.POST("/taskSubmit", v1.TaskSubmit)

		//提交mongo任务
		apiv1.POST("/taskCommonSubmit", v1.TaskCommonSubmit)

		//获取任务列表
		apiv1.GET("/getTasks", v1.GetTasks)

		//删除任务
		apiv1.GET("/deleteTask/:task_id", v1.DeleteTask)
	}

	return r
}
