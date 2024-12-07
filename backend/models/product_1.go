package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}


type Product struct {
	ID                     uint           `gorm:"primaryKey"`
	UserID                 uint           `gorm:"not null"`
	User                   User           `gorm:"foreignKey:UserID"` 
	ProductName            string         `gorm:"not null"`
	ProductDescription     string         `gorm:"not null"`
	ProductImages          gorm.JSON      `gorm:"type:json"`         
	CompressedProductImages gorm.JSON      `gorm:"type:json"`         
	ProductPrice           float64        `gorm:"not null"`
	CreatedAt              time.Time      `gorm:"autoCreateTime"`
}
