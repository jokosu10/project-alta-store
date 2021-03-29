package database

import (
	"project-alta-store/config"
	"project-alta-store/models"
)

func InsertCustomers(customer models.Customers) (interface{}, error) {
	if err := config.DB.Save(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomersByName(name string)(models.Customers,error){
	var customer models.Customers
	if e:=config.DB.Where("Username = ?",name).Find(&customer).Error; e!=nil{
		return customer,e
	}
	return customer,nil
}