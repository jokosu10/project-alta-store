package config

import (
	"fmt"
	"project-alta-store/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error reading .env file")
	}

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DBNAME"],
	)

	DB, err = gorm.Open(mysql.Open(mysqlCredentials), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitDBTest() {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error reading .env file")
	}

	mysqlCredentialsTest := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER_TEST"],
		appConfig["MYSQL_PASSWORD_TEST"],
		appConfig["MYSQL_PROTOCOL_TEST"],
		appConfig["MYSQL_HOST_TEST"],
		appConfig["MYSQL_PORT_TEST"],
		appConfig["MYSQL_DBNAME_TEST"],
	)

	fmt.Println(mysqlCredentialsTest)
	// mysqlCredentials = "root:@tcp(127.0.0.1:3306)/alterra_test?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(mysqlCredentialsTest), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&models.Categories{})
	DB.Migrator().DropTable(&models.Products{})
	DB.AutoMigrate(&models.Categories{})
	DB.AutoMigrate(&models.Products{})
}

func InitMigrate() {
	DB.AutoMigrate(&models.Customers{})
	DB.AutoMigrate(&models.Categories{})
	DB.AutoMigrate(&models.Products{})
	DB.AutoMigrate(&models.Couriers{})
	DB.AutoMigrate(&models.Orders{})
	DB.AutoMigrate(&models.Carts{})
	DB.AutoMigrate(&models.Cartitems{})
	DB.AutoMigrate(&models.Checkout_items{})
	DB.AutoMigrate(&models.Payments{})
}
