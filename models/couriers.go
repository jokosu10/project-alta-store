package models

import (
	// "gorm.io/gorm"
)

type Couriers struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"` 
	Name string `gorm:"type:varchar(255);unique" json:"name"`
}