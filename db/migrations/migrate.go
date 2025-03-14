package main

import (
	"go-backend/db"
	"go-backend/src/models"
	"log"
)

func main() {
	db.ConnectDB()

	err := db.DB.AutoMigrate(
		&models.Customer{},
		&models.Account{},
		&models.Transaction{},
		&models.AuditLog{},
		&models.LogEntry{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully")
}
