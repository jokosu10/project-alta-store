package models

import (
	// "gorm.io/gorm"
	"time"
)

type Categories struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"` 
	Name string `gorm:"type:varchar(255);unique" json:"name"`
	Description string `gorm:"type:varchar(255);" json:"description"`
	CreatedAt time.Time
	UpdatedAt time.Time

}