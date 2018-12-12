package routers

import (
	"github.com/gin-gonic/gin"
	"strings"
	"fmt"
	"net/http"
	"golang-gin-web/pkg/setting"

	"golang-gin-web/routers/api/v1"
)

func InitRouter() *gin.Engine {
	
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//上传文件
		apiv1.POST("/upload", v1.UploadFile)

		//获取数据来源
		apiv1.GET("/getDataSource", v1.GetDataSource)

		//获取品牌
		apiv1.GET("/getBrands", v1.GetBrands)

		//跑批结果文件下载
		//apiv1.GET("downFile/:task_id", v1.DownFile)

		//查看任务进度
		apiv1.GET("/taskProcess", v1.TaskProcess)

		//csv文件上传
		//apiv1.POST("/upload", v1.UploadFile)

		//提交csv任务
		apiv1.POST("/taskSubmit", v1.TaskSubmit)

		//提交mongo任务
		apiv1.POST("/taskCommonSubmit", v1.TaskCommonSubmit)

		//获取任务列表
		apiv1.GET("/getTask", v1.GetTasks)

		//删除任务
		apiv1.GET("/deleteTask", v1.DeleteTask)

		//登陆
		apiv1.POST("/login", v1.Login)

		//登出
		apiv1.GET("/logout", v1.Logout)
	}

	return r
}

//添加跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method		//请求方法
		origin := c.Request.Header.Get("Origin")		//请求头部
		var headerKeys []string								// 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")		// 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")		//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//				允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")		// 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")		// 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")		//	跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")		// 设置返回格式是json
		}
 
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()		//	处理请求
	}
}