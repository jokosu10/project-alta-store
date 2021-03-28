package models

import (
	"time"
)

type Carts struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"` 
	Customers_id int  `gorm:"type:bigint(20);foreignkey" json:"customers_id"`
	Status string `gorm:"type:varchar(255);unique" json:"status"`
	CreatedAt time.Time
}