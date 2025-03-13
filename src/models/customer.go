package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	NIK       string    `gorm:"size:16;unique;not null"`
	Phone     string    `gorm:"size:15;unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
