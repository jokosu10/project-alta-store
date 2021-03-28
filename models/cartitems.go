package models

import (

)

type Cartitems struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"` 
	carts_id int  `gorm:"type:bigint(20);foreignkey" json:"carts_id"` 
	products_id int  `gorm:"type:bigint(20);foreignkey" json:"products_id"` 
	Quantity int `json:"quantity"`
}