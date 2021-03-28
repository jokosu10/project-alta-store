package models

import (
	"time"
)

type Checkout_items struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"` 
	order_id int  `gorm:"type:bigint(20);foreignkey" json:"order_id"` 
	products_id int  `gorm:"type:bigint(20);foreignkey" json:"products_id"` 
	Quantity int `json:"quantity"`
	CreatedAt time.Time
}