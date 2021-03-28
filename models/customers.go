package models

import (
	// "gorm.io/gorm"
	"time"
)

type Customers struct {
	ID                  int    `gorm:"type:bigint(20);primarykey;AUTO_INCREMENT" json:"id"`
	Username            string `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email               string `gorm:"type:varchar(30);unique;not null" json:"email"`
	Password            string `gorm:"type:varchar(30);not null" json:"password"`
	Address             string `gorm:"type:varchar(255);" json:"address"`
	Bank_name           string `gorm:"type:varchar(255);" json:"bank_name"`
	Bank_account_number string `gorm:"type:bigint(20);" json:"bank_account_number"`
	CreatedAt           time.Time
}
