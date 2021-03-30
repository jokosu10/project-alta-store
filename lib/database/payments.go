package database

import (
	"project-alta-store/config"
	"project-alta-store/models"
	"time"
)

func InsertPayments(order models.Orders_post,order_id int) (error) {
	var newPayment models.Payments
	timeFormat := "2006-01-02 15:04:05"
	newPayment.Order_id = order_id
	newPayment.Payment_amount = order.Payment_amount
	newPayment.Payment_method = order.Payment_method
	newPayment.Payment_status = order.Payment_status
	newPayment.Payment_start_date,_= time.Parse(timeFormat,order.Payment_start_date)
	newPayment.Payment_end_date,_= time.Parse(timeFormat,order.Payment_end_date)

	if err := config.DB.Save(&newPayment).Error; err != nil {
		return err
	}
	return nil
}
