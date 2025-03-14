package models

import (
	"time"

	"gorm.io/gorm"
)

type LogEntry struct {
	ID        uint      `gorm:"primaryKey"`
	AccountNo string    `gorm:"index"`
	Message   string    `gorm:"type:text"`
	LogType   string    `gorm:"type:varchar(20)"`
	Details   string    `gorm:"type:json"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func SaveLog(db *gorm.DB, accountNo string, message string) error {
	logEntry := LogEntry{
		AccountNo: accountNo,
		Message:   message,
	}
	return db.Create(&logEntry).Error
}
