package main

import (
	"testing"
)

func TestStartDeliverySuccess(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		CourierID:      0,  // isi dengan courier_id valid
		Resi:           "", // isi dengan nomor resi
		NamaPenerima:   "", // isi dengan nama penerima
		AlamatPenerima: "", // isi dengan alamat penerima
		Status:         "", // isi dengan status awal delivery
	}

	err := service.StartDelivery(delivery)
	if err != nil {
		t.Errorf("StartDelivery failed: %v", err)
	}

	if delivery.Status != "in_delivery" {
		t.Errorf("Expected status 'in_delivery', got '%s'", delivery.Status)
	}
}

func TestStartDeliveryInvalidCourier(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		CourierID: 0,  // isi dengan courier_id tidak valid
		Resi:      "", // isi dengan nomor resi
		Status:    "", // isi dengan status delivery
	}

	err := service.StartDelivery(delivery)
	if err == nil {
		t.Error("Expected error for invalid courier_id, got nil")
	}
}

func TestStartDeliveryEmptyResi(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		CourierID: 0,  // isi dengan courier_id valid
		Resi:      "", // kosongkan nomor resi
		Status:    "", // isi dengan status delivery
	}

	err := service.StartDelivery(delivery)
	if err == nil {
		t.Error("Expected error for empty resi, got nil")
	}
}

func TestStartDeliveryNotPendingStatus(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		CourierID: 0,  // isi dengan courier_id valid
		Resi:      "", // isi dengan nomor resi
		Status:    "", // isi dengan status selain pending
	}

	err := service.StartDelivery(delivery)
	if err == nil {
		t.Error("Expected error for non-pending delivery, got nil")
	}
}

func TestCompleteDeliverySuccess(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "", // isi dengan nomor resi
		Status:         "", // isi dengan status in_delivery
		CourierID:      0,  // isi dengan courier_id valid
		AlamatPenerima: "", // isi dengan alamat penerima
	}

	err := service.CompleteDelivery(delivery)
	if err != nil {
		t.Errorf("CompleteDelivery failed: %v", err)
	}

	if delivery.Status != "delivered" {
		t.Errorf("Expected status 'delivered', got '%s'", delivery.Status)
	}

	if delivery.DeliveredAt == nil {
		t.Error("DeliveredAt should not be nil")
	}
}

func TestCompleteDeliveryNotInProgress(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:      "", // isi dengan nomor resi
		Status:    "", // isi dengan status selain in_delivery
		CourierID: 0,  // isi dengan courier_id valid
	}

	err := service.CompleteDelivery(delivery)
	if err == nil {
		t.Error("Expected error for non-in_delivery status, got nil")
	}
}

func TestCompleteDeliveryNil(t *testing.T) {
	service := NewCourierService()

	err := service.CompleteDelivery(nil)
	if err == nil {
		t.Error("Expected error for nil delivery, got nil")
	}
}

func TestGetCourierDeliveries(t *testing.T) {
	service := NewCourierService()

	deliveries := []Delivery{
		{
			CourierID: 0,  // isi dengan courier_id
			Resi:      "", // isi dengan nomor resi
			Status:    "", // isi dengan status delivery
		},
		{
			CourierID: 0,  // isi dengan courier_id berbeda
			Resi:      "", // isi dengan nomor resi
			Status:    "", // isi dengan status delivery
		},
		{
			CourierID: 0,  // isi dengan courier_id yang sama
			Resi:      "", // isi dengan nomor resi
			Status:    "", // isi dengan status delivery
		},
	}

	courierDeliveries := service.GetCourierDeliveries(deliveries, 0) // isi dengan courier_id yang dicari
	if len(courierDeliveries) != 0 { // sesuaikan jumlah expected delivery
		t.Errorf("Expected deliveries count mismatch, got %d", len(courierDeliveries))
	}
}

func TestGetCourierDeliveriesNoMatch(t *testing.T) {
	service := NewCourierService()

	deliveries := []Delivery{
		{
			CourierID: 0,  // isi dengan courier_id
			Resi:      "", // isi dengan nomor resi
			Status:    "", // isi dengan status delivery
		},
		{
			CourierID: 0,  // isi dengan courier_id
			Resi:      "", // isi dengan nomor resi
			Status:    "", // isi dengan status delivery
		},
	}

	courierDeliveries := service.GetCourierDeliveries(deliveries, 0) // isi dengan courier_id yang tidak ada
	if len(courierDeliveries) != 0 {
		t.Errorf("Expected 0 deliveries, got %d", len(courierDeliveries))
	}
}

func TestValidateDeliverySuccess(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "", // isi dengan nomor resi
		CourierID:      0,  // isi dengan courier_id valid
		NamaPenerima:   "", // isi dengan nama penerima
		AlamatPenerima: "", // isi dengan alamat penerima
	}

	err := service.ValidateDelivery(delivery)
	if err != nil {
		t.Errorf("ValidateDelivery failed: %v", err)
	}
}

func TestValidateDeliveryEmptyResi(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "", // kosongkan nomor resi
		CourierID:      0,  // isi dengan courier_id valid
		NamaPenerima:   "", // isi dengan nama penerima
		AlamatPenerima: "", // isi dengan alamat penerima
	}

	err := service.ValidateDelivery(delivery)
	if err == nil {
		t.Error("Expected error for empty resi, got nil")
	}
}

func TestValidateDeliveryInvalidCourier(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "", // isi dengan nomor resi
		CourierID:      0,  // isi dengan courier_id tidak valid
		NamaPenerima:   "", // isi dengan nama penerima
		AlamatPenerima: "", // isi dengan alamat penerima
	}

	err := service.ValidateDelivery(delivery)
	if err == nil {
		t.Error("Expected error for invalid courier_id, got nil")
	}
}

func TestValidateDeliveryMissingReceiver(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "", // isi dengan nomor resi
		CourierID:      0,  // isi dengan courier_id valid
		NamaPenerima:   "", // isi dengan nama penerima
		AlamatPenerima: "", // kosongkan alamat penerima
	}

	err := service.ValidateDelivery(delivery)
	if err == nil {
		t.Error("Expected error for empty alamat_penerima, got nil")
	}
}

func TestValidateDeliveryNil(t *testing.T) {
	service := NewCourierService()

	err := service.ValidateDelivery(nil)
	if err == nil {
		t.Error("Expected error for nil delivery, got nil")
	}
}

/* benchmark digunakan untuk mengukur performa atau kecepatan eksekusi function ketika
dijalankan berulang kali. */

func BenchmarkStartDelivery(b *testing.B) {
	service := NewCourierService()

	delivery := &Delivery{
		CourierID: 0,  // isi dengan courier_id valid
		Resi:      "", // isi dengan nomor resi
		Status:    "", // isi dengan status pending
	}

	for i := 0; i < b.N; i++ {
		service.StartDelivery(delivery)
		delivery.Status = "" // reset status delivery
	}
}

func BenchmarkValidateDelivery(b *testing.B) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "", // isi dengan nomor resi
		CourierID:      0,  // isi dengan courier_id valid
		NamaPenerima:   "", // isi dengan nama penerima
		AlamatPenerima: "", // isi dengan alamat penerima
	}

	for i := 0; i < b.N; i++ {
		service.ValidateDelivery(delivery)
	}
}
