package v1

import (

	"io/ioutil"
	"strings"
	//"time"
	//"os"
	"fmt"
	//"log"
	"encoding/csv"
	"golang-gin-web/models"
	//"net/http"
	//"github.com/gin-gonic/gin"
)

//csv文件读,插入子任务至mysql
func CsvHandle(data map[string]interface{}) {
	/*
	user_id := c.Query("user_id")
	task_type := "csv"
	file_name := c.Query("file_name")
	
	is_append := c.Query("is_append")
	number_lables := com.StrTo(c.Query("number_labels")).MustInt()
	line_numbers := com.StrTo(c.Query("line_numbers")).MustInt()
	*/
	task_column_number := data["task_column_nunmber"].(int)

	file_name := data["file_name"].(string)
	csvFile := "runtime/files/input/" + file_name
	task_id := data["task_id"].(int)
	fmt.Scanf("%s", &csvFile)
	cntb, err := ioutil.ReadFile(csvFile)
	if err != nil {
		panic(err)

	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)
	for i := 1; i<sz; i++ {
		data_sub := make(map[string]interface{})
		task_text := ss[i][task_column_number]
		data_sub["task_text"] = task_text
		data_sub["task_id"] = task_id
		data_sub["task_project_name"] = data["task_project_name"]
		data_sub["number_id"] = i
		data_sub["task_type"] = data["task_type"]

		models.AddSubTask(data_sub)
	}
}