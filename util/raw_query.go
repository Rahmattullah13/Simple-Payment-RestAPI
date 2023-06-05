package util

const (
	// LOGS HISTORIES
	ALL_HISTORIES = `SELECT * FROM logs_histories;`
	CREATE_LOG    = `INSERT INTO logs_histories(log_message) VALUES($1);`

	// USERS
	CREATE_USER = `INSERT INTO users(email, password) VALUES($1, $2) RETURNING user_id;`
	LOGIN_USER  = `SELECT * FROM users WHERE email = $1`

	// CUSTOMERS
	ALL_CUSTOMER            = `SELECT * FROM customers;`
	CREATE_CUSTOMER         = `INSERT INTO customers(user_id, name, created_at) VALUES($1, $2, $3) RETURNING customer_id;`
	READ_CUSTOMER           = `SELECT customer_id, user_id, name, balance, created_at FROM customers WHERE customer_id = $1;`
	TOPUP_CUSTOMER_BALANCE  = `UPDATE customers SET balance = $1 WHERE customer_id = $2;`
	UPDATE_CUSTOMER_BALANCE = `UPDATE customers SET balance = $1 WHERE customer_id = $2;`
	DELETE_CUSTOMER         = `DELETE FROM customers WHERE customer_id = $1;`

	// MERCHANTS
	ALL_MERCHANT            = `SELECT * FROM merchants;`
	CREATE_MERCHANT         = `INSERT INTO merchants(user_id, name, created_at) VALUES($1, $2, $3) RETURNING merchant_id;`
	READ_MERCHANT           = `SELECT merchant_id, user_id, name, balance, created_at FROM merchants WHERE merchant_id = $1;`
	UPDATE_MERCHANT_BALANCE = `UPDATE merchants SET balance = $1 WHERE merchant_id = $2;`
	DELETE_MERCHANT         = `DELETE FROM merchants WHERE merchant_id = $1;`

	// BANKS
	ALL_BANK                    = `SELECT * FROM banks;`
	CREATE_BANK                 = `INSERT INTO banks(bank_account_number, name) VALUES($1, $2) RETURNING bank_id;`
	READ_BANK                   = `SELECT bank_id, name, bank_account_number, balance FROM banks WHERE bank_id = $1;`
	READ_BANK_BY_ACCOUNT_NUMBER = `SELECT bank_id, name, bank_account_number, balance FROM banks WHERE bank_account_number = $1;`
	UPDATE_BANK_BALANCE         = `UPDATE banks SET balance = $1 WHERE bank_account_number = $2;`
	DELETE_BANK                 = `DELETE FROM banks WHERE bank_id = $1;`

	// PAYMENTS
	ALL_PAYMENT    = `SELECT * FROM payments;`
	CREATE_PAYMENT = `INSERT INTO payments(sender_id, receiver_id, amount, bank_account_number, created_at) VALUES($1, $2, $3, $4, $5) RETURNING payment_id;`
	READ_PAYMENT   = `SELECT payment_id, sender_id, receiver_id, amount, bank_account_number, created_at FROM payments WHERE payment_id = $1;`
)
