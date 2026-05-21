package main

import "errors"

// --- Tracking Functions ---

// GetTrackingStatus mengambil timeline dari database (belum diimplementasikan)
func GetTrackingStatus(resi string) *TrackingResponse {
	return nil
}

// InsertTrackingEvent memasukkan event baru dengan memvalidasi resi terlebih dahulu
func InsertTrackingEvent(
	req TrackingEvent,
	v ResiValidator,
	repo TrackingRepository,
) (TrackingEvent, error) {

	if !v.Validate(req.Resi) {
		return TrackingEvent{}, errors.New("invalid resi")
	}

	err := repo.Insert(req)

	if err != nil {
		return TrackingEvent{}, err
	}

	return req, nil
}

// --- Map/Location Functions ---

// CalculateDistance memanggil API Maps/OSM untuk hitung jarak (belum diimplementasikan)
func CalculateDistance(req DistanceRequest) *DistanceResponse {
	return nil
}

// CalculateRoute optimasi rute dari origin ke destination melewati waypoints
func CalculateRoute(req RouteRequest) *DistanceResponse {
	return nil
}

// GetCourierLocation mengambil lokasi kurir saat ini dari cache/Redis
func GetCourierLocation(courierID string) *CourierLocation {
	return nil
}
