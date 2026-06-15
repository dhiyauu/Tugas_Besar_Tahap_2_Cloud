package main

import (
	"database/sql"
)

type MySQLRepository struct {
	DB *sql.DB
}

func (r MySQLRepository) Insert(transaction Transaction) error {

	query := `
	INSERT INTO transactions
	(transaction_id, order_id, amount, metode, status, timestamp)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		transaction.TransactionID,
		transaction.OrderID,
		transaction.Amount,
		transaction.Metode,
		transaction.Status,
		transaction.Timestamp,
	)

	return err
}
