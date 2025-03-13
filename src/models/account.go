package models

import (
	"time"
)

type AccountType string

const (
	Savings  AccountType = "savings"
	Checking AccountType = "checking"
	Deposits AccountType = "deposit"
)

type Account struct {
	ID          uint        `gorm:"primaryKey"`
	CustomerID  uint        `gorm:"not null"`
	AccountNo   string      `gorm:"size:20;unique;not null"`
	AccountType AccountType `gorm:"type:account_type_enum;not null"`
	Balance     float64     `gorm:"default:0"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
}
