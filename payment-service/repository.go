package main

type PaymentRepository interface {
	Insert(transaction Transaction) error
}