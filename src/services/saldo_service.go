package services

import (
	"errors"
	"go-backend/db"
	"go-backend/logs"
	"go-backend/src/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetSaldoService(accountNo string) (float64, error) {
	var account models.Account

	err := db.DB.Where("account_no = ? AND account_type = ?", accountNo, models.Savings).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			message := "Akun tidak ditemukan atau bukan tabungan"
			logData := logrus.Fields{"account_no": accountNo}

			// Log to service
			logs.LogError(accountNo, message, logData)

			// Store in database
			logs.StoreLogEntry(accountNo, message, "WARNING", logData)

			return 0, errors.New(message)
		}

		message := "Kesalahan saat mengambil data akun"
		logData := logrus.Fields{"account_no": accountNo, "error": err.Error()}

		// Log to service
		logs.LogError(accountNo, message, logData)

		// Store in database
		logs.StoreLogEntry(accountNo, message, "ERROR", logData)

		return 0, errors.New("terjadi kesalahan pada sistem")
	}

	message := "Saldo ditemukan"
	logData := logrus.Fields{"account_no": accountNo, "balance": account.Balance}

	// Log to service
	logs.LogInfo(accountNo, message, logData)

	// Store in database
	logs.StoreLogEntry(accountNo, message, "INFO", logData)

	return account.Balance, nil
}
