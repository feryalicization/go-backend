package dto

type SaldoResponse struct {
	Saldo  float64 `json:"saldo,omitempty"`
	Remark string  `json:"remark,omitempty"`
}
