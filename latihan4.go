// Latihan 4: Slice of struct dan perulangan dengan range.
package main

// Transaksi merepresentasikan sebuah transaksi dengan ID, total, dan status.
type Transaksi struct {
	ID     int
	Total  int
	Status string
}

// Latihan4 menjalankan latihan keempat: slice of struct dan range loop.
func Latihan4() {
	// Membuat slice berisi data transaksi
	riwayatTransaksi := []Transaksi{
		{ID: 1, Total: 100000, Status: "Lunas"},
		{ID: 2, Total: 200000, Status: "Gagal"},
		{ID: 3, Total: 150000, Status: "Pending"},
	}

	// Iterasi slice dengan range, cetak hanya transaksi yang lunas
	for _, transaksi := range riwayatTransaksi {
		if transaksi.Status == "Lunas" {
			println("Transaksi ID:", transaksi.ID, "berhasil dengan total:", transaksi.Total)
		}
	}
}
