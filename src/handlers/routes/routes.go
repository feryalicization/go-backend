package routes

import (
	"go-backend/src/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")
	api.POST("/daftar", handlers.RegisterNasabah)
}
