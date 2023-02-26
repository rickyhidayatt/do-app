package utils

const (

	// TRANSACTIONS
	INSERT_BALANCE      = "INSERT INTO balance (user_id, balance) VALUES (:user_id, :balance)"
	ADD_BALANCE         = "UPDATE balances SET balance = balance + :balance WHERE user_id = :user_id"
	SEND_BALANCE        = "UPDATE balances SET balance = balance - :balance WHERE user_id = :user_id"
	CHECK_BALANCE_BY_ID = "SELECT * FROM balances WHERE user_id = $1"

	//USERS
	USER_BY_ID = "SELECT * FROM users WHERE id=$1"
)
