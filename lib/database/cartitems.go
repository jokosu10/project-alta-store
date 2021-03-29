package database

import (
	"project-alta-store/config"
	"project-alta-store/models"
	"errors"
)

func GetCartitemsByCartsId(cart_id int) ([]models.Cartitems,error){
	var cartitems [] models.Cartitems

	if e:=config.DB.Where("carts_id = ?",cart_id).Find(&cartitems).Error; e!=nil{
		return cartitems,e
	}
	
	return cartitems,nil
}

func InsertCartitems(cartitems models.Cartitems) (interface{},error){

	if err := config.DB.Save(&cartitems).Error; err!=nil{
		return nil,err
	}
	return cartitems,nil
}

func GetCartitemsById(id int) (models.Cartitems,error){
	var cartitems models.Cartitems

	if rows:=config.DB.Find(&cartitems,id).RowsAffected; rows<1{
		err := errors.New("cartitems not found")
		return cartitems,err
	}
	
	return cartitems,nil
}

func DeleteCartitemsById(id int)error{

	rows := config.DB.Delete(&models.Cartitems{},id).RowsAffected
	if rows==0{
		err := errors.New("cartitems to be deleted is not found")
		return err
	}
	return nil

}