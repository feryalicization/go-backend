package services

import (
	"errors"
	"go-backend/db"
	"go-backend/src/models"
	"log"

	"gorm.io/gorm"
)

// DepositService menambah saldo ke akun dengan nomor rekening tertentu
func DepositService(accountNo string, amount float64) (float64, error) {
	// Validasi jumlah deposit
	if amount <= 0 {
		log.Println("[ERROR] Jumlah deposit harus lebih dari 0")
		return 0, errors.New("jumlah deposit harus lebih dari 0")
	}

	var account models.Account
	err := db.DB.Where("account_no = ? AND account_type = ?", accountNo, models.Savings).First(&account).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[ERROR] Akun tidak ditemukan atau bukan jenis 'savings'")
			return 0, errors.New("akun tidak ditemukan atau bukan jenis 'savings'")
		}
		log.Println("[ERROR] Database error:", err)
		return 0, errors.New("terjadi kesalahan pada database")
	}

	// Tambahkan saldo ke akun
	account.Balance += amount

	// Simpan perubahan saldo
	if err := db.DB.Save(&account).Error; err != nil {
		log.Println("[ERROR] Gagal memperbarui saldo akun:", err)
		return 0, errors.New("gagal memperbarui saldo akun")
	}

	// Catat transaksi di tabel `transactions`
	transaction := models.Transaction{
		AccountID: account.ID,
		Type:      models.Deposit,
		Amount:    amount,
	}

	if err := db.DB.Create(&transaction).Error; err != nil {
		log.Println("[ERROR] Gagal mencatat transaksi:", err)
		return 0, errors.New("gagal mencatat transaksi")
	}

	// Catat audit log
	auditLog := models.AuditLog{
		TransactionID: transaction.ID,
		AccountID:     account.ID,
		Type:          string(models.Deposit),
		Amount:        amount,
		BalanceAfter:  account.Balance,
	}

	if err := db.DB.Create(&auditLog).Error; err != nil {
		log.Println("[ERROR] Gagal mencatat audit log:", err)
		return 0, errors.New("gagal mencatat audit log")
	}

	log.Printf("[INFO] Deposit berhasil. No Rekening: %s, Saldo Akhir: %.2f\n", accountNo, account.Balance)
	return account.Balance, nil
}
