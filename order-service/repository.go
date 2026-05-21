package main

type OrderRepository interface {
	Save(order Order) error
}