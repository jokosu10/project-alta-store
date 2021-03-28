package models

import (
	// "gorm.io/gorm"
	"time"
)

type Orders struct {
	ID           int    `gorm:"type:bigint(20);primarykey" json:"id"`
	Customers_id int    `gorm:"type:bigint(20);foreignkey;not null" json:"customers_id"`
	Couriers_id  int    `gorm:"type:bigint(20);foreignkey;not null" json:"couriers_id"`
	Province     string `gorm:"type:varchar(255);not null" json:"province"`
	City         string `gorm:"type:varchar(255);not null" json:"city"`
	Address      string `gorm:"type:varchar(255);not null" json:"address"`
	Zipcode      string `gorm:"type:varchar(255);not null" json:"zipcode"`
	CreatedAt    time.Time
}
