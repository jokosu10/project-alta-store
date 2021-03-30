package models

type Cartitems struct {
	ID          int      `gorm:"primarykey;not null;AUTO_INCREMENT" json:"id"`
	Carts_id    int      `gorm:"not null" json:"carts_id"`
	Products_id int      `gorm:"foreignkey;not null" json:"products_id"`
	Quantity    int      `gorm:"not null" json:"quantity"`
	Carts       Carts    `gorm:foreignkey:ID;references:carts_id`
	Products    Products `gorm:foreignkey:ID;references:products_id`
}

type Cartitems_Post struct {
	Carts_id    *int `json:"carts_id" validate:"required"`
	Products_id *int `json:"products_id" validate:"required"`
	Quantity    *int `json:"quantity" validate:"required"`
}

type Cartitems_Update struct {
	Quantity *int `json:"quantity" validate:"required"`
}

type Cartitems_response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []Cartitems `json:"data"`
}

type Cartitems_response_single struct {
	Code    int       `json:"code"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    Cartitems `json:"data"`
}
