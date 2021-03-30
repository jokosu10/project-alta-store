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

func GetCustomersByid(id int) (models.Customers, error) {
	var customer models.Customers
	var empty models.Customers

	if err := config.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		return empty, err
	}
	if err := config.DB.Find(&customer, id).Error; err != nil {
		return empty, err
	}

	return customer, nil
}

func GetCustomersByEmail(email string) (models.Customers, error) {
	var customer models.Customers
	var empty models.Customers

	if err := config.DB.Where("email = ?", email).First(&customer).Error; err != nil {
		return empty, err
	}
	if err := config.DB.Find(&customer, email).Error; err != nil {
		return empty, err
	}

	return customer, nil
}

func LoginCustomers(customer models.Customers) (interface{}, error) {
	var err error
	// var customerLogin models.Customers_login

	if err = config.DB.Where("email = ? AND password = ?", customer.Email, customer.Password).First(customer).Error; err != nil {
		return nil, err
	}

	// customerLogin.Token, err = middlewares.CreateToken(int(customerLogin.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
