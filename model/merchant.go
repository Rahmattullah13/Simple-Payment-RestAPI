package model

import "time"

type Merchant struct {
	MerchantId string    `json:"merchant_id" db:"merchant_id"`
	UserId     string    `json:"user_id" db:"user_id"`
	Name       string    `json:"name" db:"name"`
	Balance    int       `json:"balance" db:"balance"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type MerchantRequest struct {
	UserId string `json:"user_id" db:"user_id" example:"1"`
	Name   string `json:"name" db:"name" example:"John Doe store"`
}

type MerchantsResponse struct {
	Message string      `json:"message" example:"Get all merchants"`
	Data    interface{} `json:"data"`
}

type MerchantResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
