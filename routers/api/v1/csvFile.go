package v1

import (
	//"fmt"
	//"golang-gin-web/models"
	"golang-gin-web/pkg/e"
	"golang-gin-web/pkg/upload"
	//"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, csvFile, err := c.Request.FormFile("csvFile")
	if err != nil {
		log.Println(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}
	if csvFile == nil {
		code = e.INVALID_PARAMS
	} else {
		csvName := upload.GetCsvName(csvFile.Filename)
		//log.Println(csvName)
		fullPath := upload.GetCsvFullPath()
		savePath := upload.GetCsvPath()

		src := fullPath + csvName

		if !upload.CheckCsvExt(csvName) || !upload.CheckCsvSize(file) {
			code = e.ERROR_UPLOAD_CHECK_CSV_FORMAT
		} else {
			err := upload.CheckCsv(fullPath)
			if err != nil {
				log.Println(err)
				code = e.ERROR_UPLOAD_CHECK_CSV_FAIL

			} else if err := c.SaveUploadedFile(csvFile, src); err != nil {
				log.Println(err)
				code = e.ERROR_UPLOAD_SAVE_CSV_FAIL
			} else {
				data["csv_url"] = upload.GetCsvFullUrl(csvName)
				data["csv_save_url"] = savePath + csvName
			}
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
/*
func DownFile(c *gin.Context) {
	task_id := c.Param("task_id")
	data := make(map[string]interface{})
	data["task_id"] = task_id

	file_name, task_project_name := models.GetFileName(data)

	content, err := ioutil.ReadAll(file_name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename=%s"%file_name)
	c.Header("Content-Type", "application/text/csv")
	c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	c.Writer.Write([]byte(content))
}
*/