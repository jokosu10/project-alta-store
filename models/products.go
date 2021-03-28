package models

import (
	// "gorm.io/gorm"
	"time"

)

type Products struct{
	ID        int `gorm:"primarykey;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(255);unique;not null"`
	Categories_id int  `gorm:"type:bigint(20);foreignkey;not null"` 
	Description string `gorm:"type:varchar(255);not null"`
	Quantity int `gorm:"not null"`
	Price float32 `gorm:"type:float;not null"`
	Unit string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Products_post struct{
	Name string `json:"name" form:"name" validate:"required"`
	Categories_id int  `json:"categories_id" form:"categories_id" validate:"required"` 
	Description string `json:"description" form:"description" validate:"required"`
	Quantity int `json:"quantity" form:"quantity" validate:"required"`
	Price float32 `json:"price" form:"price" validate:"required"`
	Unit string `json:"unit" form:"unit" validate:"required"`
}

type Products_update struct{
	Name string `json:"name" form:"name"`
	Categories_id int  `json:"categories_id" form:"categories_id"` 
	Description string `json:"description" form:"description"`
	Quantity int `json:"quantity" form:"quantity"`
	Price float32 `json:"price" form:"price"`
	Unit string `json:"unit" form:"unit"`
}


type Products_response struct{
	Code string
	Message string
	Status string
	Data interface{}
}