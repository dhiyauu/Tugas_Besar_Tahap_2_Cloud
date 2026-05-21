package main

import (
	"errors"
	"time"
)

type CourierService struct {
}

func NewCourierService() *CourierService {
	return &CourierService{}
}

func (s *CourierService) StartDelivery(d *Delivery) error {

	if d == nil {
		return errors.New("delivery nil")
	}

	if d.Resi == "" {
		return errors.New("resi kosong")
	}

	d.Status = "in_delivery"

	return nil
}

func (s *CourierService) CompleteDelivery(d *Delivery) error {

	if d == nil {
		return errors.New("delivery nil")
	}

	if d.Status != "in_delivery" {
		return errors.New("delivery belum berjalan")
	}

	now := time.Now()

	d.Status = "delivered"
	d.DeliveredAt = &now

	return nil
}

func (s *CourierService) GetCourierDeliveries(
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

func (s *CourierService) ValidateDelivery(d *Delivery) error {

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
