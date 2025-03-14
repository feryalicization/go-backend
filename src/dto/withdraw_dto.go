package dto

type WithdrawRequest struct {
	NoRekening string  `json:"no_rekening" validate:"required"`
	Nominal    float64 `json:"nominal" validate:"required,gt=0"`
}

type WithdrawResponse struct {
	Saldo  float64 `json:"saldo,omitempty"`
	Remark string  `json:"remark,omitempty"`
}
