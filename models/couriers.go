package models

import (
	// "gorm.io/gorm"
)

type Couriers struct{
	ID        int `gorm:"type:bigint(20);primarykey;AUTO_INCREMENT" json:"id"` 
	Name string `gorm:"type:varchar(255);unique;not null" json:"name"`
}