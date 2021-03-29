package models

import (
	"time"
)

type Customers struct {
	ID                  int    `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Username            string `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email               string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password            string `gorm:"type:varchar(255);not null" json:"password"`
	Address             string `gorm:"type:varchar(255);" json:"address"`
	Bank_name           string `gorm:"type:varchar(255);" json:"bank_name"`
	Bank_account_number string `gorm:"type:bigint(20);default:0;" json:"bank_account_number"`
	Token               string `gorm:"type:varchar(255);" json:"token"`
	CreatedAt           time.Time
}

type Customers_register struct {
	Username string `gorm:"type:varchar(255);unique;not null" json:"username" validate:"required"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(255);not null" json:"password" validate:"required"`
}

type Customers_response struct {
	Code    string
	Message string
	Status  string
}
