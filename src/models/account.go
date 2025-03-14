package models

import (
	"time"
)

type AccountType string

const (
	Savings = "savings"
	Giro    = "giro"
)

type Account struct {
	ID          uint      `gorm:"primaryKey"`
	CustomerID  uint      `gorm:"not null"`
	AccountNo   string    `gorm:"size:20;unique;not null"`
	AccountType string    `gorm:"size:10;not null"`
	Balance     float64   `gorm:"default:0"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	// relasi table
	Customer Customer `gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE"`
}
