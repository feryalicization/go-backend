package main

import (
	"go-backend/db"
	"go-backend/src/models"
	"log"
)

func main() {
	db.ConnectDB()

	db.DB.Exec(`DO $$ BEGIN 
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'account_type_enum') THEN 
			CREATE TYPE account_type_enum AS ENUM ('savings', 'checking', 'deposit'); 
		END IF;
	END $$;`)

	db.DB.Exec(`DO $$ BEGIN 
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type_enum') THEN 
			CREATE TYPE transaction_type_enum AS ENUM ('deposit', 'withdraw'); 
		END IF;
	END $$;`)

	err := db.DB.AutoMigrate(&models.Customer{}, &models.Account{}, &models.Transaction{}, &models.AuditLog{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully")
}
