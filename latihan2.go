package main

import "fmt"

type Aplikasi struct {
	Nama  string
	Versi float64
	Fitur string
}

func upgradeFitur(aplikasi *Aplikasi) {
	aplikasi.Versi = 2.0
	aplikasi.Fitur = "Premium"
}

func Latihan2() {
	aplikasi := Aplikasi{
		Nama:  "Tracker App",
		Versi: 1.0,
		Fitur: "Basic",
	}

	fmt.Println("Data Awal")
	fmt.Println(aplikasi.Nama, aplikasi.Versi, aplikasi.Fitur)

	fmt.Println("Data Setelah Upgrade")
	upgradeFitur(&aplikasi)
	fmt.Println(aplikasi.Nama, aplikasi.Versi, aplikasi.Fitur)
}
