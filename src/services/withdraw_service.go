package services

import (
	"errors"
	"go-backend/db"
	"go-backend/src/models"
	"log"

	"gorm.io/gorm"
)

func WithdrawService(accountNo string, amount float64) (float64, error) {
	var account models.Account

	err := db.DB.Where("account_no = ? AND account_type = ?", accountNo, models.Savings).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[ERROR] No rekening tidak ditemukan atau bukan akun savings:", accountNo)
			return 0, errors.New("no rekening tidak ditemukan atau bukan akun savings")
		}
		log.Println("[ERROR] Database query error:", err)
		return 0, errors.New("terjadi kesalahan database")
	}

	// Cek apakah saldo cukup untuk penarikan
	if account.Balance < amount {
		log.Println("[ERROR] Saldo tidak cukup untuk penarikan:", account.Balance)
		return 0, errors.New("saldo tidak cukup untuk penarikan")
	}

	// Update saldo rekening
	account.Balance -= amount
	if err := db.DB.Save(&account).Error; err != nil {
		log.Println("[ERROR] Gagal memperbarui saldo:", err)
		return 0, errors.New("gagal melakukan penarikan")
	}

	// Simpan transaksi penarikan
	transaction := models.Transaction{
		AccountID: account.ID,
		Type:      models.Withdraw,
		Amount:    amount,
	}
	if err := db.DB.Create(&transaction).Error; err != nil {
		log.Println("[ERROR] Gagal menyimpan transaksi penarikan:", err)
		return 0, errors.New("gagal menyimpan transaksi")
	}

	// Simpan log audit
	auditLog := models.AuditLog{
		TransactionID: transaction.ID,
		AccountID:     account.ID,
		Type:          string(models.Withdraw),
		Amount:        amount,
		BalanceAfter:  account.Balance,
	}
	if err := db.DB.Create(&auditLog).Error; err != nil {
		log.Println("[ERROR] Gagal menyimpan audit log:", err)
		return 0, errors.New("gagal menyimpan audit log")
	}

	log.Println("[INFO] Penarikan berhasil, saldo sekarang:", account.Balance)
	return account.Balance, nil
}
