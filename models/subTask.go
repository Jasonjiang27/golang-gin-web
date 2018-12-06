package models

import (
	_ "github.com/jinzhu/gorm"
)

type SubTask struct {
	TaskId int `json:"task_id"` //任务id
	TaskUid int `json:"task_uid" gorm:"index"`//子任务id
	TaskText string `json:"task_text"` //子任务文本数据
	TaskProjectName string `json:"task_project_name"` //任务分类树名
	NumberId int 	`json:"number_id"` //任务子csv文件的第几行（专为csv类型）
	TaskType string	`json:"task_type"`	//任务类型
}

func SubTaskCount(maps interface{}) (count int) {
	
	db.Model(&SubTask{}).Where(maps).Count(&count)
	return
}

func AddSubTask 