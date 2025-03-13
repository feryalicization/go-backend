package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterCustomer(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Customer registered successfully"})
}
