package interfaces

import (
	"github.com/andrianprasetya/go-assesment-test/internal/dto/response"
	"github.com/andrianprasetya/go-assesment-test/internal/model"
)

type UserRepository interface {
	RegisterUser(user *model.User) error
	GetByNoRekening(no_rekening string) (*model.User, error)
	Update(id string, accountBalance int) error
}

type UserUsecase interface {
	RegisterUser(name, nik, no_hp string) (*response.UserRekeningResponse, error)
	GetUserByNoRekening(no_rekening string) (*response.UserBalanceResponse, error)
}
