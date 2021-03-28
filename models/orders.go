package models

import (
	// "gorm.io/gorm"
	"time"
)

type Orders struct {
	ID           int    `gorm:"type:bigint(20);primarykey;AUTO_INCREMENT" json:"id"`
	Customers_id int    `gorm:"type:bigint(20);foreignkey;not null" json:"customers_id"`
	Couriers_id  int    `gorm:"type:bigint(20);foreignkey;not null" json:"couriers_id"`
	Address      string `gorm:"type:varchar(255);not null" json:"address"`
	CreatedAt    time.Time
}
