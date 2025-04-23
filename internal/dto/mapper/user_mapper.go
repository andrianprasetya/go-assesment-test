package mapper

import (
	"github.com/andrianprasetya/go-assesment-test/internal/dto/response"
	"github.com/andrianprasetya/go-assesment-test/internal/model"
)

func FromUserModel(user *model.User) *response.UserResponse {
	return &response.UserResponse{
		Name: user.Name,
		Nik:  user.Nik,
		NoHp: user.NoHp,
	}
}

func FromUserModelBalance(balance int) *response.UserBalanceResponse {
	return &response.UserBalanceResponse{
		Balance: balance,
	}
}
