package handlers

import (
	"go-backend/src/dto"
	"go-backend/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Withdraw godoc
// @Summary Withdraw money from an account
// @Description Menarik dana dari rekening tabungan
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body dto.WithdrawRequest true "Request Body"
// @Success 200 {object} dto.WithdrawResponse
// @Failure 400 {object} dto.WithdrawResponse
// @Failure 500 {object} dto.WithdrawResponse
// @Router /tarik [post]
func WithdrawHandler(c echo.Context) error {
	var req dto.WithdrawRequest

	// Bind JSON ke struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.WithdrawResponse{Remark: "Invalid request"})
	}

	// Validasi input
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.WithdrawResponse{Remark: err.Error()})
	}

	// Panggil service untuk menarik dana
	saldo, err := services.WithdrawService(req.NoRekening, req.Nominal)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.WithdrawResponse{Remark: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.WithdrawResponse{
		Saldo:  saldo,
		Remark: "Withdrawal successful",
	})
}
