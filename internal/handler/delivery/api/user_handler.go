package api

import (
	"github.com/andrianprasetya/go-assesment-test/internal/dto/request"
	"github.com/andrianprasetya/go-assesment-test/internal/dto/response"
	"github.com/andrianprasetya/go-assesment-test/internal/dto/validation"
	"github.com/andrianprasetya/go-assesment-test/internal/interfaces"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUC interfaces.UserUsecase
}

func NewUserHandler(userUC interfaces.UserUsecase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var req request.RegisterUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ErrorResponse[any](err.Error(), nil))
	}

	if err := validation.NewValidator().Validate(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		errorMessages := validation.MapValidationErrorsToJSONTags(req, errs)
		return c.Status(fiber.StatusBadRequest).JSON(response.ValidationResponse[any](errorMessages))
	}

	user, err := h.userUC.RegisterUser(req.Name, req.Nik, req.NoHp)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse[any](err.Error(), err))
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse("User registered successfully", user))
}

func (h *UserHandler) GetBalance(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening")

	if noRekening == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no_rekening not found"})
	}

	user, err := h.userUC.GetUserByNoRekening(noRekening)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse[any](err.Error(), err))
	}

	// Jika user tidak ditemukan, kembalikan response kosong
	if user == nil {
		// Jika data tidak ditemukan, maka response kosong
		return c.Status(fiber.StatusOK).JSON(response.SuccessResponse("User not found", map[string]interface{}{}))
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse("User get successfully", user))
}
