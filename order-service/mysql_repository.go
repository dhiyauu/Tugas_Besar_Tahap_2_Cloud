package main

type MySQLRepository struct{}

func (r MySQLRepository) Save(order Order) error {

	query := `
	INSERT INTO orders
	(user_id, nama_barang, berat, dimensi,
	jenis, alamat_pengirim, alamat_penerima, status)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		order.UserID,
		order.NamaBarang,
		order.Berat,
		order.Dimensi,
		order.Jenis,
		order.AlamatPengirim,
		order.AlamatPenerima,
		order.Status,
	)

	return err
}