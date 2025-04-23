package request

type RegisterUserRequest struct {
	Name string `json:"name" validate:"required,max=100"`
	Nik  string `json:"nik" validate:"required,max=25,unique=nik:users"`
	NoHp string `json:"no_hp" validate:"required,max=15,unique=no_hp:users"`
}
