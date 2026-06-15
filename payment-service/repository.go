package main

type PaymentRepository interface {
	Save(payment Payment) error
}