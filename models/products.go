package models

import (
	// "gorm.io/gorm"
	"time"
)

type Products struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"` 
	Name string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Categories_id int  `gorm:"type:bigint(20);foreignkey;not null" json:"categories_id"` 
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Quantity int `gorm:"not null" json:"quantity"`
	Price float32 `gorm:"type:float;not null" json:"price"`
	Unit string `gorm:"type:varchar(255);not null" json:"unit"`
	CreatedAt time.Time
	UpdatedAt time.Time
}