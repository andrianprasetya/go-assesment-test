package response

type UserResponse struct {
	Name string `json:"name"`
	Nik  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type UserBalanceResponse struct {
	Balance int `json:"balance"`
}
