package database

import (
	"project-alta-store/config"
	"project-alta-store/models"
)



func InsertCourier(courier models.Couriers) (interface{},error){

	if err := config.DB.Save(&courier).Error; err!=nil{
		return nil,err
	}
	return courier,nil
}