package models

type Cartitems struct {
	ID          int `gorm:"type:bigint(20);primarykey;not null" json:"id"`
	Carts_id    int `gorm:"type:bigint(20);foreignkey;not null" json:"carts_id"`
	Products_id int `gorm:"type:bigint(20);foreignkey;not null" json:"products_id"`
	Quantity    int `gorm:"not null" json:"quantity"`
}
