package models

import (
	// "gorm.io/gorm"
	"time"
)

type Customers struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"` 
	Username string `gorm:"type:varchar(255);unique" json:"username"`
	Email string `gorm:"type:varchar(30);unique" json:"email"`
	Password string `gorm:"type:varchar(30);" json:"password"`
	Province string `gorm:"type:varchar(255);" json:"province"`
	City string `gorm:"type:varchar(255);" json:"city"`
	Address string `gorm:"type:varchar(255);" json:"address"`
	Zipcode string `gorm:"type:varchar(255);" json:"zipcode"`
	Bank_name string `gorm:"type:varchar(255);" json:"bank_name"`
	Bank_account_number string `gorm:"type:bigint(20);" json:"bank_account_number"`
	CreatedAt time.Time
}