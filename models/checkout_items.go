package models

import (
	"time"
)

type Checkout_items struct{
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"` 
	Order_id int  `gorm:"not null" json:"order_id"` 
	Products_id int  `gorm:"foreignkey;not null" json:"products_id"` 
	Quantity int `gorm:"not null" json:"quantity"`
	Order Orders `gorm:"foreignkey:ID;references:Order_id" json:"-"`
	Product Products `gorm:"foreignkey:ID;references:Products_id" json:"-"`
	CreatedAt time.Time
}

type Checkoutitem_response struct{
	Code int  `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
	Data []Checkout_items `json:"data"`
}