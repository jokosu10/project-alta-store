package config

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project-alta-store/models"
)
var DB *gorm.DB


func InitDB(){
	connectionString := "root:@tcp(127.0.0.1:3306)/alterra_store?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB,err = gorm.Open(mysql.Open(connectionString),&gorm.Config{})

	if err != nil{
		panic(err)
	}
	InitMigrate()
}

func InitMigrate(){
	DB.AutoMigrate(&models.Products{})
	DB.AutoMigrate(&models.Cartitems{})
	DB.AutoMigrate(&models.Carts{})
	DB.AutoMigrate(&models.Categories{})
	DB.AutoMigrate(&models.Couriers{})
	DB.AutoMigrate(&models.Customers{})
	DB.AutoMigrate(&models.Checkout_items{})
	DB.AutoMigrate(&models.Orders{})
	DB.AutoMigrate(&models.Payments{})
	DB.AutoMigrate(&models.Products{})
}