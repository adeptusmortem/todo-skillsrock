package models

import "time"

type Task struct {
	ID 			int			`json:"id" gorm:"primaryKey"`
	Title       string		`json:"title" gorm:"title"`
	Description string		`json:"description" gorm:"description"`
	Status 		string		`json:"status" gorm:"type:text;default:'new';check:status IN ('new', 'in_progress', 'done')"`
	Created_at 	time.Time	`json:"created_at" gorn:"created_at"`
	Updated_at 	time.Time	`json:"updated_at" gorm:"updated_at"`
}