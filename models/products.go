package models

import (
	"time"
)


type Products struct{
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Categories_id int  `json:"categories_id"` 
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Quantity int `gorm:"not null" json:"quantity"`
	Price float32 `gorm:"type:float;not null" json:"price"`
	Unit string `gorm:"type:varchar(255);not null" json:"unit"`
	Category Categories `gorm:"foreignKey:ID;references:Categories_id" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Products_post struct {
	Name          string  `json:"name" form:"name" validate:"required"`
	Categories_id int     `json:"categories_id" form:"categories_id" validate:"required"`
	Description   string  `json:"description" form:"description" validate:"required"`
	Quantity      int     `json:"quantity" form:"quantity" validate:"required"`
	Price         float32 `json:"price" form:"price" validate:"required"`
	Unit          string  `json:"unit" form:"unit" validate:"required"`
}



type Products_update struct {
	Name          string  `json:"name" form:"name"`
	Categories_id int     `json:"categories_id" form:"categories_id"`
	Description   string  `json:"description" form:"description"`
	Quantity      int     `json:"quantity" form:"quantity"`
	Price         float32 `json:"price" form:"price"`
	Unit          string  `json:"unit" form:"unit"`
}


type Products_response struct{
	Code int  `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
	Data []Products `json:"data"`
}
