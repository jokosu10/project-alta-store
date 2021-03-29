package models

import (
	"time"
)

type Checkout_items struct{
	ID        int `gorm:"type:bigint(20);primarykey;AUTO_INCREMENT89=2323fkl;cvbnmc " json:"id"` 
	Order_id int  `gorm:"type:bigint(20);foreignkey;not null" json:"order_id"` 
	Products_id int  `gorm:"type:bigint(20);foreignkey;not null" json:"products_id"` 
	Quantity int `gorm:"not null" json:"quantity"`
	CreatedAt time.Time
}