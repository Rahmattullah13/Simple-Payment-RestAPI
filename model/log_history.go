package model

type LogHistory struct {
	LogMessage string `db:"log_message"`
}

type LogHistoryResponse struct {
	Message string      `json:"message" example:"Get all history"`
	Data    interface{} `json:"data"`
}
