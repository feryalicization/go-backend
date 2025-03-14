package handlers

import (
	"go-backend/src/dto"
	"go-backend/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DepositHandler godoc
// @Summary Deposit saldo ke akun nasabah
// @Description Menyetor saldo ke akun berdasarkan nomor rekening
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body dto.DepositRequest true "Request Body"
// @Success 200 {object} dto.DepositResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /tabung [post]
func DepositHandler(c echo.Context) error {
	var req dto.DepositRequest

	// Bind JSON ke struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: "Format request tidak valid"})
	}

	// Validasi input
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: err.Error()})
	}

	// Panggil service untuk proses deposit
	balance, err := services.DepositService(req.AccountNo, req.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: err.Error()})
	}

	// Berikan response saldo terbaru
	return c.JSON(http.StatusOK, dto.DepositResponse{Balance: balance})
}
