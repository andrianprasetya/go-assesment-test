package routes

import (
	"github.com/andrianprasetya/go-assesment-test/internal/handler/delivery/api"
	"github.com/andrianprasetya/go-assesment-test/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(c *fiber.App, userUC interfaces.UserUsecase, transactionUC interfaces.TransactionUsecase) {

	userHandler := api.NewUserHandler(userUC)
	transactionHandler := api.NewTransactionHandler(transactionUC)
	api := c.Group("/api")
	v1 := api.Group("/v1")

	account := v1.Group("/account")
	account.Post("/register", userHandler.RegisterUser)
	account.Post("/saving", transactionHandler.SavingAmount)
	account.Post("/withdraw", transactionHandler.WithdrawAmount)
	account.Get("/balance/:no_rekening", userHandler.GetBalance)
}
