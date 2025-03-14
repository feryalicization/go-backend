package main

import (
	"go-backend/db"
	"go-backend/src/handlers/routes" // Import DTO untuk validasi
	"log"

	_ "go-backend/docs"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// CustomValidator untuk Echo
type CustomValidator struct {
	validator *validator.Validate
}

// Implementasi metode Validate untuk Echo
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// @title Go Backend Test
// @version 1.0
// @description Test for managing accounts & transactions service.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	log.SetFlags(0) // Disable timestamp buffering
	db.ConnectDB()

	e := echo.New()

	// Daftarkan Validator ke Echo
	e.Validator = &CustomValidator{validator: validator.New()}

	// Swagger Reoute
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Insert All Route
	routes.RegisterRoutes(e)

	log.Println("Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
