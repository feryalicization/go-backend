package main

import (
	"go-backend/db"
	"go-backend/src/handlers/routes"
	"log"

	_ "go-backend/docs"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// CustomValidator for Echo
type CustomValidator struct {
	validator *validator.Validate
}

// Validate method for Echo
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

	// Register Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// CORS Middleware
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Swagger Route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Register All Routes
	routes.RegisterRoutes(e)

	log.Println("Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
