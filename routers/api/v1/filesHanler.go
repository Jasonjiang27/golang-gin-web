package v1

import (
	"fmt"
	//"golang-gin-web/pkg/e"
	"net/http"
	//"strings"

	"github.com/gin-gonic/gin"
)


//文件下载
func DownFile(c *gin.Context) {
	//task_id := com.StrTo(c.Query("task_id")).MustInt()

	file_out := c.Query("file_out")

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename=file_out")
	c.Header("Content-Type", "application/text/csv")
	c.Header("Accept-Length", fmt.Sprintf("%d", len(file_out)))
	c.Writer.Write([]byte(file_out))

}
