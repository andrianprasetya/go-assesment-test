package api

import (
	"github.com/andrianprasetya/go-assesment-test/internal/dto/request"
	"github.com/andrianprasetya/go-assesment-test/internal/dto/response"
	"github.com/andrianprasetya/go-assesment-test/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionUC interfaces.TransactionUsecase
}

func NewTransactionHandler(transactionUC interfaces.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{transactionUC: transactionUC}
}

func (h *TransactionHandler) SavingAmount(c *fiber.Ctx) error {
	var req request.SavingRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse[any](err.Error(), nil))
	}

	transaction, err := h.transactionUC.Create(req.NoRekening, "C", req.Amount)

	if transaction == nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse[any](err.Error(), nil))
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse[any](err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse("Transaction saving successfully", transaction))
}

func (h *TransactionHandler) WithdrawAmount(c *fiber.Ctx) error {
	var req request.SavingRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse[any](err.Error(), nil))
	}

	transaction, err := h.transactionUC.Create(req.NoRekening, "D", req.Amount)

	if transaction == nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse[any](err.Error(), nil))
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse[any](err.Error(), err))
	}
	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse("Transaction withdraw successfully", transaction))
}
