package model

import "time"

type Customer struct {
	CustomerId string    `json:"customer_id" db:"customer_id"`
	UserId     string    `json:"user_id" db:"user_id"`
	Name       string    `json:"name" db:"name"`
	Balance    int       `json:"balance" db:"balance"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type CustomerRequest struct {
	UserId string `json:"user_id" db:"user_id" example:"1"`
	Name   string `json:"name" db:"name" example:"John Doe"`
}

type CustomersResponse struct {
	Message string      `json:"message" example:"Get all customers"`
	Data    interface{} `json:"data"`
}

type CustomerResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TopUpRequest struct {
	CustomerId string `json:"customer_id" db:"customer_id" example:"1"`
	Balance    int    `json:"balance" db:"balance" example:"50000"`
}
