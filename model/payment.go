package model

import "time"

type Payment struct {
	PaymentId         string    `json:"payment_id" db:"payment_id"`
	SenderId          string    `json:"sender_id" db:"sender_id"`
	ReceiverId        string    `json:"receiver_id" db:"receiver_id"`
	Amount            int       `json:"amount" db:"amount"`
	BankAccountNumber string    `json:"bank_account_number" db:"bank_account_number"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
}

type PaymentRequest struct {
	SenderId          string `json:"sender_id" db:"sender_id" example:"1"`
	ReceiverId        string `json:"receiver_id" db:"receiver_id" example:"1"`
	Amount            int    `json:"amount" db:"amount" example:"7000"`
	BankAccountNumber string `json:"bank_account_number" db:"bank_account_number" example:"12345678"`
}

type PaymentsResponse struct {
	Message string      `json:"message" example:"Get all payments"`
	Data    interface{} `json:"data"`
}

type PaymentResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
