package database

import (
	"project-alta-store/config"
	"project-alta-store/middlewares"
	"project-alta-store/models"
)

func InsertCustomers(customer models.Customers) (interface{}, error) {
	if err := config.DB.Save(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func LoginCustomers(customer models.Customers) (interface{}, error) {
	var err error

	if err = config.DB.Where("email = ? AND password = ?", customer.Email, customer.Password).First(customer).Error; err != nil {
		return nil, err
	}

	customer.Token, err = middlewares.CreateToken(int(customer.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
