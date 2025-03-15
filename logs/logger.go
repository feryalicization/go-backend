package logs

import (
	"encoding/json"
	"go-backend/db"
	"go-backend/src/models"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger adalah instance global dari logrus
var Logger = logrus.New()

func init() {
	// Konfigurasi logger
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)
}

func LogInfo(accountNo, message string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	if accountNo != "" {
		fields["account_no"] = accountNo
	}

	Logger.WithFields(fields).Info(message)
}

func LogError(accountNo, message string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	if accountNo != "" {
		fields["account_no"] = accountNo
	}

	Logger.WithFields(fields).Error(message)
}

func StoreLogEntry(accountNo, message, logType string, fields logrus.Fields) {
	detailsJSON, _ := json.Marshal(fields)

	logEntry := models.LogEntry{
		AccountNo: accountNo,
		Message:   message,
		LogType:   logType,
		Details:   string(detailsJSON),
	}

	if err := db.DB.Create(&logEntry).Error; err != nil {
		Logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Gagal menyimpan log ke database")
	}
}
