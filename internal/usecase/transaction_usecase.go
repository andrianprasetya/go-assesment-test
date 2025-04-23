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

type transactionUsecase struct {
	transactionRepo interfaces.TransactionRepository
	userRepo        interfaces.UserRepository
}

func NewTransactionUsecase(transactionRepo interfaces.TransactionRepository, UserRepo interfaces.UserRepository) interfaces.TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		userRepo:        UserRepo,
	}
}

func (t transactionUsecase) Create(no_rekening string, type_transaction string, amount int) (*response.UserBalanceResponse, error) {
	user, errGetReq := t.userRepo.GetByNoRekening(no_rekening)
	if user == nil {
		log.WithFields(log.Fields{
			"no_rekening": no_rekening,
			"error":       "account not found",
		}).Error("account " + no_rekening + " not found")
		return nil, fmt.Errorf("account not found")
	}

	if errGetReq != nil {
		log.WithFields(log.Fields{
			"no_rekening": no_rekening,
			"error":       errGetReq,
		}).Error("failed to get user by no_rekening")
		return nil, fmt.Errorf("failed to get user by no rekening: %w", errGetReq)
	}
	var accountBalance int
	var errorMessageLogs, successMessageLogs string
	switch type_transaction {
	case "C":
		accountBalance = user.Balance + amount
		errorMessageLogs = "failed saving amount to rekening"
		successMessageLogs = "saving amount to rekening success"
	case "D":

		if user.Balance < amount {
			return nil, fmt.Errorf("account balance insufficient : %d", user.Balance)
		}
		accountBalance = user.Balance - amount
		errorMessageLogs = "failed withdraw amount to rekening"
		successMessageLogs = "withdraw amount to rekening success"
	default:
		accountBalance = 0
	}

	transaction := &model.Transaction{
		ID:              utils.GenerateID(),
		UserId:          user.ID,
		TypeTransaction: type_transaction,
		Amount:          amount,
	}

	errTransaction := t.transactionRepo.Create(transaction)
	if errTransaction != nil {
		log.WithFields(log.Fields{
			"error": errTransaction,
		}).Error("failed to create transaction")
		return nil, fmt.Errorf(errorMessageLogs)
	}
	errUser := t.userRepo.Update(user.ID, accountBalance)
	if errUser != nil {
		log.WithFields(log.Fields{
			"id":          user.ID,
			"no_rekening": no_rekening,
			"error":       errUser,
		}).Error("failed to update balance")
		return nil, fmt.Errorf(errorMessageLogs+" : %w", errUser)
	}

	log.WithFields(log.Fields{
		"id":          user.ID,
		"no_rekening": no_rekening,
		"amount":      amount,
	}).Info(successMessageLogs)
	return mapper.FromUserModelBalance(accountBalance), errTransaction
}
