// internal/repository/question.go
package repository

import (
	"github.com/joshua468/access-bank-exam-system/models"
	"gorm.io/gorm"
)

func CreateQuestion(db *gorm.DB, question *models.Question) error {
	return db.Create(question).Error
}

func GetQuestionsByExamID(db *gorm.DB, examID uint) ([]models.Question, error) {
	var questions []models.Question
	err := db.Where("exam_id = ?", examID).Find(&questions).Error
	return questions, err
}
