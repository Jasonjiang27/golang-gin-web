package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Task struct {
	//Model

	task_id 			 int 	`json:"task_id" gorm:"index"` 	//任务id

	user_id           int    `json:"user_id"`            	//创建任务的用户id
	task_status       string `json:"task_status"`        	//任务状态
	task_type         string `json:"task_type"`          	//来源是csv还是数据库
	file_name		 string `json:"file_name"`          	//对于csv文件名
	file_location     string `json:"file_location"`      	//文件位置
	task_projectName  string `json:"task_project_name"`  	//分类树名
	task_columnNumber int    `json:"task_column_number"` 	//分类数据列名

	data_source		 string	`json:"data_source"`			//数据来源（mongo数据源）
	limit            int    `json:"limit"`              	//条数限制
	start_time        string `json:"start_time"`         	//开始时间
	end_time          string `json:"end_time"`           	//结束时间
	sub_task_numbers   int    `json:"sub_task_numbers"`   	//子任务数

}

func GetDataSource(data string) (dataSource []string) {
	//补充mongo数据库来源
	return
}

func GetBrands(name string, series []string) (data map[string][]map[string]interface{}){
	//补充mongo查询的数据
	return
}

func GetTasks(pageNum int, pageSize int, maps interface{}) (tasks []Task) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tasks)

	return
}

func GetTasksTotal(maps interface{}) (count int) {
	db.Model(&Task{}).Where(maps).Count(&count)

	return
}

//判断任务是否存在
func ExistTaskById(task_id int) bool {
	var task Task
	db.Select("task_id").Where("task_id = ?", task_id).First(&task)

	if task.task_id > 0 {
		return true
	}
	return false
}

func TaskSubmit(data map[string]interface{}) error {
	task := Task{
		//TaskId: data["task_id"].(int),
		user_id: data["user_id"].(int),
		task_type: data["task_type"].(string),
		file_name: data["file_name"].(string),
		file_location: data["file_location"].(string),
		task_status: data["task_status"].(string),
		task_projectName: data["task_project_name"].(string),
		task_columnNumber: data["task_column_number"].(int),
		limit: data["limit"].(int),
		start_time: data["start_time"].(string),
		end_time: data["end_time"].(string),
		sub_task_numbers: data["sub_task_numbers"].(int),
	}
	if err := db.Create(&task).Error; err != nil {
		return err
	}
	
		return nil
	}

func TaskCommonSubmit(data map[string]interface{}) error {
	task := Task{
		user_id: data["user_id"].(int),
		task_type: data["task_type"].(string),
		task_status: data["task_status"].(string),
		task_projectName: data["task_project_name"].(string),
		task_columnNumber: data["task_column_number"].(int),
		data_source: data["data_source"].(string),
		limit: data["limit"].(int),
		start_time: data["start_time"].(string),
		end_time: data["end_time"].(string),
		sub_task_numbers: data["sub_task_numbers"].(int),
	}
	if err := db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTask(task_id int) bool {
	db.Where("task_id = ?", task_id).Delete(Task{})

	return true
}

func (task *Task) BeforeTask(scope *gorm.Scope) error {
	scope.SetColumn("start_time", time.Now().Unix())

	return nil
}

func (task *Task) AfterTask(scope *gorm.Scope) error {
	scope.SetColumn("end_time", time.Now().Unix())

	return nil
}