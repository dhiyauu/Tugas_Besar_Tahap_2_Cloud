package main

func CreatePayment(
	req Payment,
	repo PaymentRepository,
) (Payment, error) {

	req.Status = "pending"

	err := repo.Save(req)

	if err != nil {
		return Payment{}, err
	}

	return req, nil
}

func GetPayment(id int) *Payment {
	return nil
}

func UpdatePaymentStatus(id int, status string) bool {
	return false
}