package models

import (

	"time"
)

type Carts struct{
	ID        int `gorm:"type:bigint(20);primarykey;AUTO_INCREMENT" json:"id"` 
	Customers_id int  `gorm:"type:bigint(20);foreignkey;not null" json:"customers_id"`
	Status string `gorm:"type:varchar(255);unique;not null" json:"status"`
	CartItem []Cartitems `gorm:"foreignKey:Carts_id"`
	CreatedAt time.Time
}