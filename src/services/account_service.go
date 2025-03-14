package services

import (
	"errors"
	"fmt"
	"go-backend/db"
	"go-backend/logs"
	"go-backend/src/models"
	"math/rand"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RegisterNasabahService(name, nik, phone, accountType string) (string, error) {
	// Validasi jenis akun menggunakan constant dari models
	validAccountTypes := map[string]bool{
		models.Savings: true,
		models.Giro:    true,
	}

	if !validAccountTypes[accountType] {
		message := "Jenis akun tidak valid"
		logData := logrus.Fields{"account_type": accountType}

		logs.LogError(nik, message, logData)
		logs.StoreLogEntry(nik, message, "WARNING", logData)

		return "", errors.New("jenis akun tidak valid, gunakan: savings atau giro")
	}

	// Cari Customer berdasarkan NIK
	var existingCustomer models.Customer
	err := db.DB.Where("nik = ?", nik).First(&existingCustomer).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		message := "Database query failed"
		logData := logrus.Fields{"error": err.Error(), "nik": nik}

		logs.LogError(nik, message, logData)
		logs.StoreLogEntry(nik, message, "ERROR", logData)

		return "", errors.New("database error")
	}

	// Jika NIK sudah ada, periksa apakah sudah memiliki akun dengan jenis yang sama
	if err == nil {
		var existingAccount models.Account
		err = db.DB.Where("customer_id = ? AND account_type = ?", existingCustomer.ID, accountType).First(&existingAccount).Error
		if err == nil {
			message := "Customer sudah memiliki akun dengan jenis yang sama"
			logData := logrus.Fields{"nik": nik, "account_type": accountType}

			logs.LogInfo(nik, message, logData)
			logs.StoreLogEntry(nik, message, "WARNING", logData)

			return "", errors.New(message)
		}

		// Cek apakah nomor HP sudah digunakan oleh NIK lain
		var phoneCheck models.Customer
		err = db.DB.Where("phone = ? AND nik != ?", phone, nik).First(&phoneCheck).Error
		if err == nil {
			message := "No HP sudah digunakan oleh pengguna lain"
			logData := logrus.Fields{"nik": nik, "phone": phone}

			logs.LogInfo(nik, message, logData)
			logs.StoreLogEntry(nik, message, "WARNING", logData)

			return "", errors.New(message)
		}

		// Buat akun baru dengan jenis akun yang diinginkan
		newAccount := models.Account{
			CustomerID:  existingCustomer.ID,
			AccountNo:   GenerateAccountNumber(),
			AccountType: accountType,
			Balance:     0.0,
		}

		if err := db.DB.Create(&newAccount).Error; err != nil {
			message := "Gagal registrasi akun"
			logData := logrus.Fields{"nik": nik, "account_type": accountType, "error": err.Error()}

			logs.LogError(nik, message, logData)
			logs.StoreLogEntry(nik, message, "ERROR", logData)

			return "", errors.New(message)
		}

		message := "Akun berhasil dibuat"
		logData := logrus.Fields{"nik": nik, "account_no": newAccount.AccountNo, "account_type": accountType}

		logs.LogInfo(nik, message, logData)
		logs.StoreLogEntry(nik, message, "INFO", logData)

		return newAccount.AccountNo, nil
	}

	// Jika NIK belum ada, buat customer baru
	newCustomer := models.Customer{
		Name:  name,
		NIK:   nik,
		Phone: phone,
	}

	if err := db.DB.Create(&newCustomer).Error; err != nil {
		message := "Gagal menyimpan data nasabah"
		logData := logrus.Fields{"nik": nik, "error": err.Error()}

		logs.LogError(nik, message, logData)
		logs.StoreLogEntry(nik, message, "ERROR", logData)

		return "", errors.New(message)
	}

	// Buat akun pertama untuk customer baru
	newAccount := models.Account{
		CustomerID:  newCustomer.ID,
		AccountNo:   GenerateAccountNumber(),
		AccountType: accountType,
		Balance:     0.0,
	}

	if err := db.DB.Create(&newAccount).Error; err != nil {
		message := "Gagal registrasi akun"
		logData := logrus.Fields{"nik": nik, "account_type": accountType, "error": err.Error()}

		logs.LogError(nik, message, logData)
		logs.StoreLogEntry(nik, message, "ERROR", logData)

		return "", errors.New(message)
	}

	message := "Nasabah baru dan akun berhasil dibuat"
	logData := logrus.Fields{
		"nik":          nik,
		"name":         name,
		"phone":        phone,
		"account_no":   newAccount.AccountNo,
		"account_type": accountType,
	}

	logs.LogInfo(nik, message, logData)
	logs.StoreLogEntry(nik, message, "INFO", logData)

	return newAccount.AccountNo, nil
}

// GenerateAccountNumber generates a unique account number
func GenerateAccountNumber() string {
	return "98" + fmt.Sprintf("%08d", rand.Intn(100000000))
}
