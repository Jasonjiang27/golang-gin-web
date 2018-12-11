package v1

import (
	"strconv"
	"io/ioutil"
	"strings"
	"time"
	//"log"
	"encoding/csv"
	"golang-gin-web/models"
	"golang-gin-web/pkg/util"
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
	task_column_number := data["task_column_number"].(int)
	task_project_name := data["task_project_name"].(string)
	file_name := data["file_name"].(string)
	user_id := data["user_id"].(int)
	task_status := data["task_status"]
	task_type := data["task_type"]
	csvFile := "runtime/files/input/" + file_name
	
	//fmt.Scanf("%s", &csvFile)
	cntb, err := ioutil.ReadFile(csvFile)
	
	if err != nil {
		panic(err)

	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)
	start_time := time.Now().Unix()
	tm := time.Unix(start_time, 0)
	start_time_str := tm.Format("2006-01-02 03:04:05 PM")

	//插入总任务表
	data_total := make(map[string]interface{})
	data_total["task_id"] = util.EncodeMD5(file_name + task_project_name + strconv.Itoa(user_id))
	data_total["sub_task_numbers"] = sz
	data_total["start_time"] = start_time_str
	data_total["file_name"] = file_name
	data_total["file_location"] = csvFile
	data_total["user_id"] = user_id
	data_total["task_status"] = task_status
	data_total["task_type"] = task_type
	data_total["task_project_name"] = task_project_name
	data_total["task_column_number"] = task_column_number
	//data_total["limit"] = data["limit"]
	models.TaskSubmit(data_total)
	

	for i := 1; i<sz; i++ {
		//插入子任务表
		data_sub := make(map[string]interface{})
		var task_text string
		task_text = ss[i][task_column_number]

		data_sub["task_id"] = util.EncodeMD5(file_name + task_project_name + strconv.Itoa(user_id))
		data_sub["task_text"] = task_text
		data_sub["task_project_name"] = task_project_name
		data_sub["number_id"] = i
		data_sub["task_type"] = task_type

		models.AddSubTask(data_sub)
	}
}