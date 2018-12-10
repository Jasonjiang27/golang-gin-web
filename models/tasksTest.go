package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	//Model

	TaskId 			string `json:"task_id" gorm:"index"` //任务id

	UserId           int    `json:"user_id"`            //创建任务的用户id
	TaskStatus       string `json:"task_status"`        //任务状态
	TaskType         string `json:"task_type"`          //来源是csv还是数据库
	FileName         string `json:"file_name"`          //对于csv文件名
	FileLocation     string `json:"file_location"`      //文件位置
	TaskProjectName  string `json:"task_project_name"`  //分类树名
	TaskColumnNumber int    `json:"task_column_number"` //分类数据列名

	DataSource     string `json:"data_source"`      //数据来源（mongo数据源）
	Limit          int    `json:"limit"`            //条数限制
	StartTime      string `json:"start_time"`       //开始时间
	EndTime        string `json:"end_time"`         //结束时间
	SubTaskNumbers int    `json:"sub_task_numbers"` //子任务数

}

func GetDataSource(data string) (dataSource []string) {
	//补充mongo数据库来源
	return
}

func GetBrands(name string, series []string) (data map[string][]map[string]interface{}) {
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
/*
//判断任务是否存在
func ExistTaskById(taskId string) bool {
	var task Task
	
	return db.Select("taskId").Where("taskId = ?", taskId).First(&task)
}
*/

//提交csv任务
func TaskSubmit(data map[string]interface{}) error {
	task := Task{
		TaskId:	 		  data["task_id"].(string),
		UserId:           data["user_id"].(int),
		TaskType:         data["task_type"].(string),
		FileName:         data["file_name"].(string),
		FileLocation:     data["file_location"].(string),
		TaskStatus:       data["task_status"].(string),
		TaskProjectName:  data["task_project_name"].(string),
		TaskColumnNumber: data["task_column_number"].(int),
		//Limit:            data["limit"].(int),
		StartTime:        data["start_time"].(string),
		//EndTime:          data["end_time"].(string),
		SubTaskNumbers:   data["sub_task_numbers"].(int),
	}
	if err := db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}

//提交舆情任务
func TaskCommonSubmit(data map[string]interface{}) error {
	task := Task{
		UserId:          data["user_id"].(int),
		TaskType:        data["task_type"].(string),
		TaskStatus:      data["task_status"].(string),
		TaskProjectName: data["task_project_name"].(string),

		DataSource:     data["data_source"].(string),
		Limit:          data["limit"].(int),
		StartTime:      data["start_time"].(string),
		EndTime:        data["end_time"].(string),
		SubTaskNumbers: data["sub_task_numbers"].(int),
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
