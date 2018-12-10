package v1
/*
import (
	//"github.com/Unknwon/com"
	"fmt"
	"strings"
	"net/http"
	"golang-gin-web/pkg/e"
	"github.com/gin-gonic/gin"

)

//上传csv文件
func UploadFile(c *gin.Context) {
	//name := c.PostForm("name")
	//fmt.Println(name)
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	filename := header.Filename
	fmt.Println(file, err, filename)
	file_format := strings.Split(filename, ".")
	if file_format[len(file_format)-1] != "csv" {
		c.String(http.StatusBadRequest, "请上传csv格式文件")
	} else {
		code := e.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"msg":  "上传文件成功",
			"code": code,
		})
	}

}

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
*/
