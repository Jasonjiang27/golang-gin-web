package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"golang-gin-web/pkg/file"
	"golang-gin-web/pkg/setting"
)

func GetCsvFullUrl(name string) string {
	return setting.AppSetting.CsvPrefixUrl + "/" + GetCsvPath() + name
}

func GetCsvName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)

	return fileName + ext
}

func GetCsvPath() string {
	return setting.AppSetting.CsvSavePath
}

func GetCsvFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetCsvPath()
}

func CheckCsvExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.CsvAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckCsvSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}
	return size <= setting.AppSetting.CsvMaxSize
}

func CheckCsv(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
