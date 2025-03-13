package models

import "time"

type AuditLog struct {
	ID            uint      `gorm:"primaryKey"`
	TransactionID uint      `gorm:"not null"`
	AccountID     uint      `gorm:"not null"`
	Type          string    `gorm:"type:transaction_type_enum;not null"`
	Amount        float64   `gorm:"not null"`
	BalanceAfter  float64   `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}
