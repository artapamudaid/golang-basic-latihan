package main

type Transaksi struct {
	ID     int
	Total  int
	Status string
}

func Latihan4() {
	riwayatTransaksi := []Transaksi{
		{ID: 1, Total: 100000, Status: "Lunas"},
		{ID: 2, Total: 200000, Status: "Gagal"},
		{ID: 3, Total: 150000, Status: "Pending"},
	}

	for _, transaksi := range riwayatTransaksi {
		if transaksi.Status == "Lunas" {
			println("Transaksi ID:", transaksi.ID, "berhasil dengan total:", transaksi.Total)
		}
	}
}
