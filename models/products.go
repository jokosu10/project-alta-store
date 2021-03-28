package models

import (
	// "gorm.io/gorm"
	"time"
)

type Products struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"` 
	Name string `gorm:"type:varchar(255);unique" json:"name"`
	Categories_id int  `gorm:"type:bigint(20);foreignkey" json:"categories_id"` 
	Description string `gorm:"type:varchar(255)" json:"description"`
	Quantity int `json:"quantity"`
	Price float32 `gorm:"type:float" json:"price"`
	Unit string `gorm:"type:varchar(255)" json:"unit"`
	CreatedAt time.Time
	UpdatedAt time.Time
}