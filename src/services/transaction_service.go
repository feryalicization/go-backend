package services

import (
	"errors"
	"go-backend/db"
	"go-backend/logs"
	"go-backend/src/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// DepositService menambah saldo ke akun dengan nomor rekening tertentu
func DepositService(accountNo string, amount float64) (float64, error) {
	// Validasi jumlah deposit
	if amount <= 0 {
		message := "Jumlah deposit harus lebih dari 0"
		logData := logrus.Fields{"account_no": accountNo, "amount": amount}

		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "WARNING", logData)

		return 0, errors.New(message)
	}

	var account models.Account
	err := db.DB.Where("account_no = ? AND account_type = ?", accountNo, models.Savings).First(&account).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			message := "Akun tidak ditemukan atau bukan jenis 'savings'"
			logData := logrus.Fields{"account_no": accountNo}

			logs.LogError(accountNo, message, logData)
			logs.StoreLogEntry(accountNo, message, "WARNING", logData)

			return 0, errors.New(message)
		}

		message := "Database error saat mencari akun"
		logData := logrus.Fields{"account_no": accountNo, "error": err.Error()}

		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "ERROR", logData)

		return 0, errors.New("terjadi kesalahan pada database")
	}

	// Tambahkan saldo ke akun
	account.Balance += amount

	// Simpan perubahan saldo
	if err := db.DB.Save(&account).Error; err != nil {
		message := "Gagal memperbarui saldo akun"
		logData := logrus.Fields{"account_no": accountNo, "amount": amount, "error": err.Error()}

		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "ERROR", logData)

		return 0, errors.New(message)
	}

	// Catat transaksi di tabel `transactions`
	transaction := models.Transaction{
		AccountID: account.ID,
		Type:      models.Deposit,
		Amount:    amount,
	}

	if err := db.DB.Create(&transaction).Error; err != nil {
		message := "Gagal mencatat transaksi"
		logData := logrus.Fields{"account_no": accountNo, "amount": amount, "error": err.Error()}

		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "ERROR", logData)

		return 0, errors.New(message)
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
		message := "Gagal mencatat audit log"
		logData := logrus.Fields{"account_no": accountNo, "amount": amount, "balance_after": account.Balance, "error": err.Error()}

		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "ERROR", logData)

		return 0, errors.New(message)
	}

	// Log sukses
	message := "Deposit berhasil"
	logData := logrus.Fields{
		"account_no":    accountNo,
		"amount":        amount,
		"balance_after": account.Balance,
	}

	logs.LogInfo(accountNo, message, logData)
	logs.StoreLogEntry(accountNo, message, "INFO", logData)

	return account.Balance, nil
}
