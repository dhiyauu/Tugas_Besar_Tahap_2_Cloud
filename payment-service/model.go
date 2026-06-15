package main

type Payment struct {
	PaymentID int    `json:"payment_id"`
	OrderID   int    `json:"order_id"`
	Method    string `json:"method"`
	Option    string `json:"option"`
	Status    string `json:"status"`
}