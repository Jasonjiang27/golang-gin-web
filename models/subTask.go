package models

import (
	_ "github.com/jinzhu/gorm"
)

//子任务
type SubTasks struct {
	TaskId 			int 	`json:"task_id"` //任务id
	TaskUid 		int 	`json:"task_uid" gorm:"index"`//子任务id
	TaskText 		string 	`json:"task_text"` //子任务文本数据
	TaskProjectName string 	`json:"task_project_name"` //任务分类树名
	NumberId 		int 	`json:"number_id"` //任务子csv文件的第几行（专为csv类型）
	TaskType 		string	`json:"task_type"`	//任务类型
	Task			Task	`json:"task"` 	//任务结构体
}

//子任务结果
type TaskResult struct {
	TaskUid 		int 	`json:"task_uid"` //子任务id
	NumberId 		int 	`json:"number_id"` //csv文件中的行数
	Result 			string 	`json:"result"` //分词结果
	TaskText 		string 	`json:"task_text"` //任务的文本数据
	Status 			string 	`json:"status"` //是否分词成功（成功、失败等）
	AnswerJudge 	string 	`json:"answer_judge"` //人工判断是否成功
}

//获取子任务总数
func SubTaskCount(maps interface{}) (count int) {
	
	db.Model(&SubTasks{}).Where(maps).Count(&count)
	return
}

//获取子任务完成数
func SubTaskDoneCount(maps interface{}) (count int) {

	db.Model(&TaskResult{}).Where(maps).Count(&count)
	return
}

//创建子任务插入数据表
func AddSubTask(data map[string]interface{}) error {
	subTask := SubTasks{
		TaskId:				data["task_id"].(int),
		TaskText:			data["task_text"].(string),
		TaskProjectName:	data["task_project_name"].(string),
		NumberId:			data["number_id"].(int),
		TaskType:			data["task_type"].(string),
	}

	if err := db.Preload("Task").Create(&subTask).Error; err != nil {
		return err
	}
	return nil
}