package models

import "time"

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)

type Transaction struct {
	ID        uint            `gorm:"primaryKey"`
	AccountID uint            `gorm:"not null"`
	Type      TransactionType `gorm:"type:transaction_type_enum;not null"`
	Amount    float64         `gorm:"not null;check:amount>0"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
}
