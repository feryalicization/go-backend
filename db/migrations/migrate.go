package main

import (
	"go-backend/db"
	"go-backend/src/models"
	"log"
)

func main() {
	db.ConnectDB()

	// 🛠️ Pastikan ENUM sudah dibuat sebelum AutoMigrate()
	db.DB.Exec(`DO $$ BEGIN 
	    CREATE TYPE account_type_enum AS ENUM ('savings', 'checking', 'deposit'); 
	EXCEPTION WHEN duplicate_object THEN null; END $$;`)

	db.DB.Exec(`DO $$ BEGIN 
	    CREATE TYPE transaction_type_enum AS ENUM ('deposit', 'withdraw'); 
	EXCEPTION WHEN duplicate_object THEN null; END $$;`)

	// 🛠️ Jalankan AutoMigrate
	err := db.DB.AutoMigrate(
		&models.Customer{},
		&models.Account{},
		&models.Transaction{},
		&models.AuditLog{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully")
}
