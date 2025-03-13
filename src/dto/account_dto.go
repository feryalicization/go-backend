package dto

type RegisterRequest struct {
	Nama        string `json:"nama" validate:"required"`
	NIK         string `json:"nik" validate:"required,len=16"`
	NoHP        string `json:"no_hp" validate:"required"`
	AccountType string `json:"account_type" validate:"required,oneof=savings checking deposit"`
}

type RegisterResponse struct {
	NoRekening string `json:"no_rekening,omitempty"`
	Remark     string `json:"remark,omitempty"`
}
