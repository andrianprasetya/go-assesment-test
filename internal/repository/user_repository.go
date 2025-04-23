package repository

import (
	"errors"
	"github.com/andrianprasetya/go-assesment-test/internal/interfaces"
	"github.com/andrianprasetya/go-assesment-test/internal/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{DB: db}
}

func (r userRepository) RegisterUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r userRepository) GetByID(id string) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, "id = ?", id).Error
	return &user, err
}

func (r userRepository) GetByNoRekening(no_rekening string) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, "no_rekening = ?", no_rekening).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Data Not Found, Return nil and error nil
		return nil, nil
	}

	return &user, err
}

func (r userRepository) Update(id string, accountBalance int) error {
	return r.DB.Model(&model.User{}).Where("id = ?", id).Update("balance", accountBalance).Error
}
