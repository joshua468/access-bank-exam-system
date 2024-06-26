package models

import (
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	UserID uint `json:"user_id" gorm:"not null"`
	ExamID uint `json:"exam_id" gorm:"not null"`
	Score  int  `json:"score" gorm:"not null"`
}
