package models

import (

	"time"
)

type Carts struct{
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"` 
	Customers_id int  `gorm:"not null" json:"customers_id"`
	Status string `gorm:"type:varchar(255);unique;not null" json:"status"`
	Customer Customers `gorm:"foreignKey:ID;references:Customers_id"`
	CreatedAt time.Time
}

type Carts_update struct{
	Status string `json:"status" validate:"required"`
}