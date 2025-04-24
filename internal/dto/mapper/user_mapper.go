package mapper

import (
	"github.com/andrianprasetya/go-assesment-test/internal/dto/response"
)

func FromUserModelBalance(balance int) *response.UserBalanceResponse {
	return &response.UserBalanceResponse{
		Balance: balance,
	}
}

func FromUserModelRekening(no_rekening string) *response.UserRekeningResponse {
	return &response.UserRekeningResponse{
		NoRekening: no_rekening,
	}
}
