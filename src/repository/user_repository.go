package repository

import (
	"api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByFirebaseUID(user *model.User, firebaseUID string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByFirebaseUID(user *model.User, firebaseUID string) error {
	if err := ur.db.Where("firebase_uid=?", firebaseUID).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
