package repository

import (
	"github.com/joshua468/access-bank-exam-system/models"
	"gorm.io/gorm"
)

func CreateExam(db *gorm.DB, exam *models.Exam) error {
	return db.Create(exam).Error
}

func ListExams(db *gorm.DB) ([]models.Exam, error) {
	var exams []models.Exam
	err := db.Preload("Questions").Find(&exams).Error
	return exams, err
}
