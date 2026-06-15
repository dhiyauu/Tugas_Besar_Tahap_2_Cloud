package main

import "testing"

type MockPaymentRepository struct{}

func (m MockPaymentRepository) Save(
	payment Payment,
) error {
	return nil
}

func TestCreatePayment(t *testing.T) {

	mockRepo := MockPaymentRepository{}

	req := Payment{
		OrderID: 1,
		Method:  "E-Wallet",
		Option:  "DANA",
	}

	p, err := CreatePayment(
		req,
		mockRepo,
	)

	if err != nil {
		t.Fatal(err)
	}

	if p.Status != "pending" {
		t.Errorf(
			"expected status pending but got %s",
			p.Status,
		)
	}
}
