package main

import (
	"errors"
	"time"
)

// StartDelivery memulai pengiriman
func StartDelivery(d *Delivery) error {

	if d == nil {
		return errors.New("delivery nil")
	}

	d.Status = "in_delivery"

	return nil
}

// CompleteDelivery menyelesaikan pengiriman
func CompleteDelivery(d *Delivery) error {

	if d == nil {
		return errors.New("delivery nil")
	}

	now := time.Now()

	d.Status = "delivered"
	d.DeliveredAt = &now

	return nil
}

// GetCourierDeliveries mengambil delivery berdasarkan courier
func GetCourierDeliveries(
	deliveries []Delivery,
	courierID int,
) []Delivery {

	var result []Delivery

	for _, d := range deliveries {

		if d.CourierID == courierID {
			result = append(result, d)
		}
	}

	return result
}

// ValidateDelivery validasi sederhana
func ValidateDelivery(d *Delivery) error {

	if d == nil {
		return errors.New("delivery nil")
	}

	if d.Resi == "" {
		return errors.New("resi kosong")
	}

	if d.CourierID <= 0 {
		return errors.New("courier invalid")
	}

	return nil
}
