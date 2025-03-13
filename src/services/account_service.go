package services

import (
	"errors"
	"fmt"
	"go-backend/db"
	"go-backend/src/models"
	"log"
	"math/rand"

	"gorm.io/gorm"
)

func RegisterNasabahService(name, nik, phone, accountType string) (string, error) {
	// Validasi jenis akun
	validTypes := map[string]models.AccountType{
		"savings":  models.Savings,
		"checking": models.Checking,
		"deposit":  models.Deposits,
	}

	accType, exists := validTypes[accountType]
	if !exists {
		log.Println("[ERROR] Jenis akun tidak valid:", accountType)
		return "", errors.New("jenis akun tidak valid, gunakan: savings, checking, atau deposit")
	}

	// Cari Customer berdasarkan NIK
	var existingCustomer models.Customer
	err := db.DB.Where("nik = ?", nik).First(&existingCustomer).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("[ERROR] Database query failed:", err)
		return "", errors.New("database error")
	}

	// Jika NIK sudah ada, periksa apakah sudah memiliki akun dengan jenis yang sama
	if err == nil {
		var existingAccount models.Account
		err = db.DB.Where("customer_id = ? AND account_type = ?", existingCustomer.ID, accType).First(&existingAccount).Error
		if err == nil {
			log.Println("[INFO] Customer dengan NIK ini sudah memiliki akun jenis:", accountType)
			return "", errors.New("customer dengan NIK ini sudah memiliki akun jenis " + accountType)
		}

		// Cek apakah nomor HP sudah digunakan oleh NIK lain
		var phoneCheck models.Customer
		err = db.DB.Where("phone = ? AND nik != ?", phone, nik).First(&phoneCheck).Error
		if err == nil {
			log.Println("[INFO] No HP sudah digunakan oleh NIK lain:", phone)
			return "", errors.New("No HP sudah digunakan oleh pengguna lain")
		}

		// Buat akun baru dengan jenis akun yang diinginkan
		newAccount := models.Account{
			CustomerID:  existingCustomer.ID,
			AccountNo:   GenerateAccountNumber(),
			AccountType: accType,
			Balance:     0.0,
		}

		log.Printf("[DEBUG] Menyimpan akun baru untuk customer: %+v\n", newAccount)
		if err := db.DB.Create(&newAccount).Error; err != nil {
			log.Println("[ERROR] Gagal registrasi akun:", err)
			return "", errors.New("gagal registrasi akun")
		}

		log.Println("[INFO] Akun berhasil dibuat dengan No Rekening:", newAccount.AccountNo)
		return newAccount.AccountNo, nil
	}

	// Jika NIK belum ada, buat customer baru
	newCustomer := models.Customer{
		Name:  name,
		NIK:   nik,
		Phone: phone,
	}

	log.Printf("[DEBUG] Menyimpan nasabah baru: %+v\n", newCustomer)
	if err := db.DB.Create(&newCustomer).Error; err != nil {
		log.Println("[ERROR] Gagal menyimpan data nasabah:", err)
		return "", errors.New("gagal menyimpan data nasabah")
	}

	log.Println("[INFO] Nasabah berhasil disimpan dengan ID:", newCustomer.ID)

	// save to db
	newAccount := models.Account{
		CustomerID:  newCustomer.ID,
		AccountNo:   GenerateAccountNumber(),
		AccountType: accType,
		Balance:     0.0,
	}

	log.Printf("[DEBUG] Menyimpan akun pertama: %+v\n", newAccount)
	if err := db.DB.Create(&newAccount).Error; err != nil {
		log.Println("[ERROR] Gagal registrasi akun:", err)
		return "", errors.New("gagal registrasi akun")
	}

	log.Println("[INFO] Akun berhasil dibuat dengan No Rekening:", newAccount.AccountNo)
	return newAccount.AccountNo, nil
}

// GenerateAccountNumber generates a unique account number
func GenerateAccountNumber() string {
	return "98" + fmt.Sprintf("%08d", rand.Intn(100000000))
}
