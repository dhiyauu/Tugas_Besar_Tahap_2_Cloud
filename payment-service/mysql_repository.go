package main

type MySQLRepository struct{}

func (r MySQLRepository) Save(payment Payment) error {

	query := `
	INSERT INTO payments
	(order_id, method, payment_option, status)
	VALUES (?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		payment.OrderID,
		payment.Method,
		payment.Option,
		payment.Status,
	)

	return err
}