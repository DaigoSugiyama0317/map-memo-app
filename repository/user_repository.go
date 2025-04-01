package repository

import (
	"github.com/DaigoSugiyama0317/map-memo-app/model"
	"gorm.io/gorm"
)

type IUser_Repository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	UpdatePassword(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

//コンストラクタ関数
func NewUserRepository(db *gorm.DB) IUser_Repository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("Email=?", email).First(user).Error; err != nil {
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

//後で実装
func (ur *userRepository) UpdatePassword(user *model.User) error {
	return nil
}