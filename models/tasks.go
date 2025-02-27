package models

import (
	"time"
)

type Task struct {
	ID 			int			`json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string		`json:"title" gorm:"type:text;not null"`
	Description string		`json:"description" gorm:"type:text"`
	Status 		string		`json:"status" gorm:"type:text;default:'new';check:status IN ('new', 'in_progress', 'done')"`
	Created_at 	time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at 	time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}