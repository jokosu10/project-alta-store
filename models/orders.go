package models

import (
	// "gorm.io/gorm"
	"time"
)

type Orders struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"`
	Customers_id int  `gorm:"type:bigint(20);foreignkey" json:"customers_id"` 
	Couriers_id int  `gorm:"type:bigint(20);foreignkey" json:"couriers_id"` 
	Province string `gorm:"type:varchar(255);" json:"province"`
	City string `gorm:"type:varchar(255);" json:"city"`
	Address string `gorm:"type:varchar(255);" json:"address"`
	Zipcode string `gorm:"type:varchar(255);" json:"zipcode"`
	CreatedAt time.Time
}
