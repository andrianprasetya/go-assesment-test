package interfaces

import (
	"github.com/andrianprasetya/go-assesment-test/internal/dto/response"
	"github.com/andrianprasetya/go-assesment-test/internal/model"
)

type TransactionRepository interface {
	Create(user *model.Transaction) error
}

type TransactionUsecase interface {
	Create(no_rekening string, type_transaction string, amount int) (*response.UserBalanceResponse, error)
}
