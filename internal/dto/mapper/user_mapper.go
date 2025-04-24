package mapper

import (
	"github.com/andrianprasetya/go-assesment-test/internal/dto/response"
)

func FromUserModelBalance(balance int) *response.UserBalanceResponse {
	return &response.UserBalanceResponse{
		Balance: balance,
	}
}
