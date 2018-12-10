package api

import (
	"golang-gin-web/pkg/e"
	"golang-gin-web/pkg/upload"
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