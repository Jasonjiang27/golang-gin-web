package util

import (
	"golang-gin-web/pkg/setting"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

//获取页码和总数，在setting中每页10条，从第2页开始加载更多的数据
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
