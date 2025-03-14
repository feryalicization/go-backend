package services

import (
	"errors"
	"go-backend/db"
	"go-backend/logs"
	"go-backend/src/models"

	"gorm.io/gorm"
)

func WithdrawService(accountNo string, amount float64) (float64, error) {
	// Validasi jumlah penarikan
	if amount <= 0 {
		message := "Jumlah penarikan harus lebih dari 0"
		logData := map[string]interface{}{
			"accountNo": accountNo,
			"amount":    amount,
		}
		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "WARNING", logData)
		return 0, errors.New(message)
	}

	var account models.Account

	// Cari akun berdasarkan nomor rekening dan tipe akun savings
	err := db.DB.Where("account_no = ? AND account_type = ?", accountNo, models.Savings).First(&account).Error
	if err != nil {
		message := "No rekening tidak ditemukan atau bukan akun savings"
		logData := map[string]interface{}{
			"accountNo": accountNo,
			"error":     err.Error(),
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.LogError(accountNo, message, logData)
			logs.StoreLogEntry(accountNo, message, "WARNING", logData)
			return 0, errors.New(message)
		}

		message = "Terjadi kesalahan database"
		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "ERROR", logData)
		return 0, errors.New(message)
	}

	// Periksa apakah saldo cukup untuk penarikan
	if account.Balance < amount {
		message := "Saldo tidak cukup untuk penarikan"
		logData := map[string]interface{}{
			"accountNo":      accountNo,
			"currentBalance": account.Balance,
			"withdrawAmount": amount,
		}
		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "CRITICAL", logData)
		return 0, errors.New(message)
	}

	// Kurangi saldo rekening
	account.Balance -= amount
	if err := db.DB.Save(&account).Error; err != nil {
		message := "Gagal memperbarui saldo rekening"
		logData := map[string]interface{}{
			"accountNo": accountNo,
			"error":     err.Error(),
		}
		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "ERROR", logData)
		return 0, errors.New(message)
	}

	// Simpan transaksi penarikan
	transaction := models.Transaction{
		AccountID: account.ID,
		Type:      models.Withdraw,
		Amount:    amount,
	}
	if err := db.DB.Create(&transaction).Error; err != nil {
		message := "Gagal mencatat transaksi penarikan"
		logData := map[string]interface{}{
			"accountNo": accountNo,
			"error":     err.Error(),
		}
		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "ERROR", logData)
		return 0, errors.New(message)
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
		message := "Gagal mencatat audit log"
		logData := map[string]interface{}{
			"accountNo": accountNo,
			"error":     err.Error(),
		}
		logs.LogError(accountNo, message, logData)
		logs.StoreLogEntry(accountNo, message, "ERROR", logData)
		return 0, errors.New(message)
	}

	// Logging sukses
	message := "Penarikan berhasil"
	logData := map[string]interface{}{
		"accountNo":      accountNo,
		"finalBalance":   account.Balance,
		"withdrawAmount": amount,
	}
	logs.StoreLogEntry(accountNo, message, "INFO", logData)

	return account.Balance, nil
}
