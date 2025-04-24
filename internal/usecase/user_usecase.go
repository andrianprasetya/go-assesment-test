package usecase

import (
	"fmt"
	"github.com/andrianprasetya/go-assesment-test/internal/dto/mapper"
	"github.com/andrianprasetya/go-assesment-test/internal/dto/response"
	"github.com/andrianprasetya/go-assesment-test/internal/interfaces"
	"github.com/andrianprasetya/go-assesment-test/internal/model"
	"github.com/andrianprasetya/go-assesment-test/internal/utils"
	log "github.com/sirupsen/logrus"
)

type userUsecase struct {
	UserRepo interfaces.UserRepository
}

func NewUserUsecase(userRepo interfaces.UserRepository) interfaces.UserUsecase {
	return &userUsecase{UserRepo: userRepo}
}

func (u userUsecase) RegisterUser(name, nik, no_hp string) (*response.UserBalanceResponse, error) {
	user := &model.User{
		ID:         utils.GenerateID(),
		Name:       name,
		Nik:        nik,
		NoHp:       no_hp,
		NoRekening: utils.GenerateUniqueNumber(),
	}

	err := u.UserRepo.RegisterUser(user)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("failed to create account")
		return &response.UserBalanceResponse{}, fmt.Errorf("something Went wrong")
	}

	log.WithFields(log.Fields{
		"id":          user.ID,
		"name":        user.Name,
		"nik":         user.Nik,
		"no_hp":       user.NoHp,
		"no_rekening": user.NoRekening,
	}).Info("create account success")

	return mapper.FromUserModelBalance(user.Balance), err
}

func (u userUsecase) GetUserByNoRekening(no_rekening string) (*response.UserBalanceResponse, error) {
	userRepo, err := u.UserRepo.GetByNoRekening(no_rekening)
	if userRepo == nil {
		return nil, nil
	}
	return mapper.FromUserModelBalance(userRepo.Balance), err
}
