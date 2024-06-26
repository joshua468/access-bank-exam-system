package repository

import (
	"github.com/joshua468/access-bank-exam-system/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
    return db.Create(user).Error
}

func GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
    var user models.User
    err := db.Where("username = ?", username).First(&user).Error
    return &user, err
}
