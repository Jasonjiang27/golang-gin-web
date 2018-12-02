package util

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"

	"golang-gin-web/pkg/setting"
)

//获取页码和总数，在setting中每页10条，从第2页开始加载更多的数据
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}