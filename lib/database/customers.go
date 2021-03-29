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
