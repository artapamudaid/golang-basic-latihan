// Latihan 3: Struct, error handling, dan multi-return.
package main

import (
	"errors"
	"fmt"
)

// Pembayaran merepresentasikan data hasil pembayaran.
type Pembayaran struct {
	Total  int
	Status string
}

// prosesBayar memproses pembayaran, mengembalikan Pembayaran dan error jika saldo tidak cukup.
func prosesBayar(saldo int, tagihan int) (Pembayaran, error) {
	if saldo < tagihan {
		return Pembayaran{}, errors.New("Saldo tidak cukup")
	}

	return Pembayaran{
		Total:  tagihan,
		Status: "Lunas",
	}, nil
}

// Latihan3 menjalankan latihan ketiga: error handling dan multi-return.
func Latihan3() {
	var saldo int = 500000
	var tagihan int = 450000

	// Panggil prosesBayar dan cek error
	pembayaran, err := prosesBayar(saldo, tagihan)

	if err != nil {
		fmt.Println("Gagal: ", err.Error())
		return
	}

	fmt.Println("Pembayaran Berhasil:", pembayaran.Total, pembayaran.Status)

}
