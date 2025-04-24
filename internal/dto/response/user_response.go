package response

type UserRekeningResponse struct {
	NoRekening string `json:"no_rekening"`
}

type UserBalanceResponse struct {
	Balance int `json:"balance"`
}
