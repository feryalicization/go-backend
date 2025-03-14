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
	Logger.SetFormatter(&logrus.JSONFormatter{}) // Format log ke JSON
	Logger.SetOutput(os.Stdout)                  // Output log ke terminal
	Logger.SetLevel(logrus.InfoLevel)            // Level default INFO
}

// LogInfo hanya mencetak log ke service (tidak menyimpan ke DB)
func LogInfo(accountNo, message string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	if accountNo != "" {
		fields["account_no"] = accountNo
	}

	Logger.WithFields(fields).Info(message)
}

// LogError hanya mencetak log ke service (tidak menyimpan ke DB)
func LogError(accountNo, message string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	if accountNo != "" {
		fields["account_no"] = accountNo
	}

	Logger.WithFields(fields).Error(message)
}

// storeLogEntry menyimpan log ke database secara terpisah
func StoreLogEntry(accountNo, message, logType string, fields logrus.Fields) {
	// Konversi fields menjadi JSON
	detailsJSON, _ := json.Marshal(fields)

	logEntry := models.LogEntry{
		AccountNo: accountNo,
		Message:   message,
		LogType:   logType,
		Details:   string(detailsJSON),
	}

	// Simpan ke database
	if err := db.DB.Create(&logEntry).Error; err != nil {
		Logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Gagal menyimpan log ke database")
	}
}
