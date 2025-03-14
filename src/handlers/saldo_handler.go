package handlers

import (
	"go-backend/src/dto"
	"go-backend/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetSaldo godoc
// @Summary Get account balance
// @Description Melihat saldo nasabah berdasarkan nomor rekening
// @Tags accounts
// @Accept json
// @Produce json
// @Param no_rekening path string true "Nomor Rekening"
// @Success 200 {object} dto.SaldoResponse
// @Failure 400 {object} dto.SaldoResponse
// @Failure 500 {object} dto.SaldoResponse
// @Router /saldo/{no_rekening} [get]
func SaldoHandler(c echo.Context) error {
	noRekening := c.Param("no_rekening")
	saldo, err := services.GetSaldoService(noRekening)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.SaldoResponse{Remark: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SaldoResponse{
		Saldo:  saldo,
		Remark: "Saldo retrieved successfully",
	})
}
