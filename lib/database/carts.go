package database

import (
	"project-alta-store/config"
	"project-alta-store/models"
	"errors"

)

func InsertCarts(cart models.Carts) (interface{},error){

	if err := config.DB.Save(&cart).Error; err!=nil{
		return nil,err
	}
	return cart,nil
}

func GetCartsById(id int) (models.Carts,error){
	var carts models.Carts

	if rows:=config.DB.Find(&carts,id).RowsAffected; rows<1{
		err := errors.New("carts not found")
		return carts,err
	}
	
	return carts,nil
}