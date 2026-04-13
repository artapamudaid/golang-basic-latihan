// Latihan 1: Variabel, tipe data, kondisi (if-else), perulangan (for), dan fungsi.
package main

import "fmt"

// Latihan1 menjalankan latihan pertama: variabel, kondisi, perulangan, dan fungsi.
func Latihan1() {
	// Deklarasi variabel dengan tipe data berbeda
	var namaProject string = "GO App"
	var tahunRilis int = 2022
	var isSaas bool = true

	// Percabangan: cek apakah project lama atau baru
	if tahunRilis < 2024 {
		fmt.Println("Project Lama", namaProject)
	} else {
		fmt.Println("Project Baru", namaProject)
	}

	// Perulangan 1 sampai 10, cetak apakah angka genap atau ganjil
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "(Genap)")
		} else {
			fmt.Println(i)
		}
	}

	// Panggil fungsi cekStatus untuk menentukan status aplikasi
	status := cekStatus(isSaas)

	fmt.Println(status)
}

// cekStatus mengembalikan string berdasarkan status aplikasi.
func cekStatus(appStatus bool) string {
	if appStatus {
		return "Aplikasi SaaS"
	} else {
		return "Aplikasi Project"
	}
}
