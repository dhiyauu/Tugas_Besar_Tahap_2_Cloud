package main

import "testing"

type MockValidator struct{}

func (m MockValidator) CheckUser(userID int, token string) bool {
	return true
}

type MockOrderRepository struct{}

func (m MockOrderRepository) Save(order Order) error {
	return nil
}

func TestCreateOrder(t *testing.T) {

	mockValidator := MockValidator{}
	mockRepo := MockOrderRepository{}

	req := Order{
		UserID:         1,
		NamaBarang:     "Laptop",
		Berat:          2,
		Dimensi:        "10x10",
		Jenis:          "Elektronik",
		AlamatPengirim: "Bandung",
		AlamatPenerima: "Jakarta",
	}

	o, err := CreateOrder(
		req,
		"dummy",
		mockValidator,
		mockRepo,
	)

	if err != nil {
		t.Fatal(err)
	}

	if o.Status != "created" {
		t.Errorf(
			"expected status created but got %s",
			o.Status,
		)
	}
}