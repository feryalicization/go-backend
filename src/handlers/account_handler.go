package handlers

import (
	"go-backend/src/dto"
	"go-backend/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterNasabah godoc
// @Summary Register new nasabah
// @Description Mendaftarkan nasabah baru dengan nama, NIK, no_hp, dan tipe akun
// @Tags accounts
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Request Body"
// @Success 200 {object} dto.RegisterResponse
// @Failure 400 {object} dto.RegisterResponse
// @Failure 500 {object} dto.RegisterResponse
// @Router /daftar [post]
func RegisterNasabah(c echo.Context) error {
	var req dto.RegisterRequest

	// Bind JSON ke struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.RegisterResponse{Remark: "Invalid request"})
	}

	// Validasi input
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.RegisterResponse{Remark: err.Error()})
	}

	// Panggil service dengan tambahan parameter accountType
	noRekening, err := services.RegisterNasabahService(req.Nama, req.NIK, req.NoHP, req.AccountType)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.RegisterResponse{Remark: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.RegisterResponse{
		NoRekening: noRekening,
		Remark:     "Registration successful",
	})
}
