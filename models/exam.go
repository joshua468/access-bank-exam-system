package models

import (
	"gorm.io/gorm"
)

type Exam struct {
	gorm.Model
	Title     string     `json:"title" gorm:"not null"`
	Questions []Question `json:"questions" gorm:"foreignKey:ExamID"`
}
