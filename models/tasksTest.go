package models

type Task struct {
	Model

	TaskId int `json:"task_id" gorm:"index"` //任务id

	UserId           int    `json:"user_id"`            //创建任务的用户id
	TaskStatus       int    `json:"task_status"`        //任务状态
	Type             string `json:"type"`               //来源是csv还是数据库
	TaskProjectName  string `json:"task_project_name"`  //分类树名
	TaskColumnNumber int    `json:"task_column_number"` //分类数据列名
	Limit            int    `json:"limit"`              //条数限制
	StartTime        string `json:"start_time"`         //开始时间
	EndTime          string `json:"end_time"`           //结束时间
	SubTaskNumbers   int    `json:"sub_task_numbers"`   //子任务数

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

	if task.TaskId > 0 {
		return true
	}
	return false
}

func AddTask(data map[string]interface{}) error {
	task := Task{
		TaskId: data["task_id"].(int),
		UserId: data["user_id"].(int),
		Type: data["type"].(string),
		TaskStatus: data["task_status"].(string),
		TaskProjectName: data["task_project_name"].(string),
		TaskColumnNumber: data["task_column_number"].(int),
		StartTime: data["start_time"].(string),
		EndTime: data["end_time"].(string),
		SubTaskNumbers: data["sub_task_numbers"].(int),
	}
	if err := db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTask(taskId int) error {
	if err := db.Where("taskId = ?", taskId).Delete(Task{}).Error; err != nil {
		return err
	}
	return nil
}
