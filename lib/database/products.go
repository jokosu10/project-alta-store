package database

import(
	"project-alta-store/config"
	"project-alta-store/models"
)

func GetProducts() (interface{},error){
	var product []models.Products

	if err:=config.DB.Find(&product).Error; err!=nil{
		return nil,err
	}
	return product,nil
}

func InsertProducts(product models.Products) (interface{},error){

	if err := config.DB.Save(&product).Error; err!=nil{
		return nil,err
	}
	return product,nil
}