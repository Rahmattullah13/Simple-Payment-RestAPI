package model

type Bank struct {
	BankId            string `json:"bank_id" db:"bank_id"`
	Name              string `json:"name" db:"name"`
	BankAccountNumber string `json:"bank_account_number" db:"bank_account_number"`
	Balance           int    `json:"balance" db:"balance"`
}

type BankRequest struct {
	BankAccountNumber string `json:"bank_account_number" db:"bank_account_number" example:"12345678"`
	Name              string `json:"name" db:"name" example:"Bank BCAX"`
}

type BanksResponse struct {
	Message string      `json:"message" example:"Get all banks"`
	Data    interface{} `json:"data"`
}

type BankResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
