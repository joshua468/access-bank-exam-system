package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ExamID uint   `json:"exam_id" gorm:"not null"`
	Text   string `json:"text" gorm:"not null"`
	Answer string `json:"answer" gorm:"not null"`
}
