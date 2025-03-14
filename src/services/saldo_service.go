package services

import (
	"errors"
	"go-backend/db"
	"go-backend/src/models"
	"log"

	"gorm.io/gorm"
)

func GetSaldoService(accountNo string) (float64, error) {
	var account models.Account

	err := db.DB.Where("account_no = ? AND account_type = ?", accountNo, models.Savings).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[ERROR] No rekening tidak ditemukan atau bukan akun tabungan:", accountNo)
			return 0, errors.New("no rekening tidak ditemukan atau bukan akun tabungan")
		}
		log.Println("[ERROR] Database query error:", err)
		return 0, errors.New("terjadi kesalahan database")
	}

	log.Println("[INFO] Saldo ditemukan untuk no rekening:", accountNo, "Saldo:", account.Balance)
	return account.Balance, nil
}
