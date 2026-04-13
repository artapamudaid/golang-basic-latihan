// Latihan 2: Struct dan pointer — demonstrasi passing by reference.
package main

import "fmt"

// Aplikasi merepresentasikan sebuah aplikasi dengan nama, versi, dan fitur.
type Aplikasi struct {
	Nama  string
	Versi float64
	Fitur string
}

// upgradeFitur mengubah versi dan fitur aplikasi menggunakan pointer.
func upgradeFitur(aplikasi *Aplikasi) {
	aplikasi.Versi = 2.0
	aplikasi.Fitur = "Premium"
}

// Latihan2 menjalankan latihan kedua: struct dan pointer.
func Latihan2() {
	// Inisialisasi struct Aplikasi
	aplikasi := Aplikasi{
		Nama:  "Tracker App",
		Versi: 1.0,
		Fitur: "Basic",
	}

	fmt.Println("Data Awal")
	fmt.Println(aplikasi.Nama, aplikasi.Versi, aplikasi.Fitur)

	// Panggil upgradeFitur dengan passing pointer (&) agar data asli berubah
	fmt.Println("Data Setelah Upgrade")
	upgradeFitur(&aplikasi)
	fmt.Println(aplikasi.Nama, aplikasi.Versi, aplikasi.Fitur)
}
