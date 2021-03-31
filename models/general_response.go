package models

type ErrorResponse struct{
	Code int  `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
}