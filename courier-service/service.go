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

// StartDelivery - Mulai proses delivery
func (s *CourierService) StartDelivery(delivery *Delivery) error {

	if delivery == nil {
		return errors.New("delivery nil")
	}

	if delivery.Resi == "" {
		return errors.New("resi kosong")
	}

	if delivery.CourierID <= 0 {
		return errors.New("courier invalid")
	}

	if delivery.Status != "pending" {
		return errors.New("status invalid")
	}

	delivery.Status = "in_delivery"

	return nil
}

// CompleteDelivery - Selesaikan proses delivery
func (s *CourierService) CompleteDelivery(delivery *Delivery) error {

	if delivery == nil {
		return errors.New("delivery nil")
	}

	if delivery.Status != "in_delivery" {
		return errors.New("delivery belum berjalan")
	}

	now := time.Now()

	delivery.Status = "delivered"
	delivery.DeliveredAt = &now

	return nil
}

// GetCourierDeliveries - Ambil semua delivery untuk courier tertentu
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

// ValidateDelivery - Validasi delivery data
func (s *CourierService) ValidateDelivery(
	delivery *Delivery,
) error {

	if delivery == nil {
		return errors.New("delivery nil")
	}

	if delivery.Resi == "" {
		return errors.New("resi kosong")
	}

	if delivery.CourierID <= 0 {
		return errors.New("courier invalid")
	}

	if delivery.NamaPenerima == "" {
		return errors.New("nama penerima kosong")
	}

	if delivery.AlamatPenerima == "" {
		return errors.New("alamat penerima kosong")
	}

	return nil
}
