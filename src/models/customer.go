package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null" json:"nama"`
	NIK       string    `gorm:"size:16;unique;not null" json:"nik"`
	Phone     string    `gorm:"size:15;not null" json:"no_hp"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
