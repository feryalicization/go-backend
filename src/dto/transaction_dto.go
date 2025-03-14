package dto

type DepositRequest struct {
	AccountNo string  `json:"no_rekening" binding:"required"`
	Amount    float64 `json:"nominal" binding:"required,gt=0"`
}

type DepositResponse struct {
	Balance float64 `json:"saldo"`
}

type ErrorResponse struct {
	Remark string `json:"remark"`
}
