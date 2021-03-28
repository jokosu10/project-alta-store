package models

import (
	"time"
)

type Payments struct{
	ID        int `gorm:"type:bigint(20);primarykey" json:"id"` 
	Order_id int  `gorm:"type:bigint(20);foreignkey" json:"order_id"` 
	Payment_method string `gorm:"type:varchar(255);" json:"payment_method"`
	Payment_start_date time.Time `json:"payment_start_date"`
	Payment_end_date time.Time `json:"payment_end_date"`
	Payment_status string `gorm:"type:varchar(255);" json:"payment_status"`
	Payment_amount float32 `gorm:"type:float" json:"payment_amount"`
}