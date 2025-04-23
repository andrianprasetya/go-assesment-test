package request

type SavingRequest struct {
	NoRekening string `json:"no_rekening" validate:"required"`
	Amount     int    `json:"amount" validate:"required,number"`
}

type WithdrawRequest struct {
	NoRekening string `json:"no_rekening" validate:"required"`
	Amount     string `json:"amount" validate:"required,number"`
}
