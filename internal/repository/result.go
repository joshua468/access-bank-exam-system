package repository

import (
	"github.com/joshua468/access-bank-exam-system/models"
	"gorm.io/gorm"
)

func CreateResult(db *gorm.DB, result *models.Result) error {
	return db.Create(result).Error
}

func GetResultsByUserID(db *gorm.DB, userID uint) ([]models.Result, error) {
	var results []models.Result
	err := db.Where("user_id = ?", userID).Find(&results).Error
	return results, err
}
