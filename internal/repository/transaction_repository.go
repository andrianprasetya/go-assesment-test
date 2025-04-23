package repository

import (
	"github.com/andrianprasetya/go-assesment-test/internal/interfaces"
	"github.com/andrianprasetya/go-assesment-test/internal/model"
	"gorm.io/gorm"
)

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) interfaces.TransactionRepository {
	return &transactionRepository{DB: db}
}

func (r transactionRepository) Create(transaction *model.Transaction) error {
	return r.DB.Create(transaction).Error
}
