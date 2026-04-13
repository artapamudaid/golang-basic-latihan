package main

import (
	"errors"
	"fmt"
)

type Pembayaran struct {
	Total  int
	Status string
}

func prosesBayar(saldo int, tagihan int) (Pembayaran, error) {
	if saldo < tagihan {
		return Pembayaran{}, errors.New("Saldo tidak cukup")
	}

	return Pembayaran{
		Total:  tagihan,
		Status: "Lunas",
	}, nil
}

func Latihan3() {
	var saldo int = 500000
	var tagihan int = 450000

	pembayaran, err := prosesBayar(saldo, tagihan)

	if err != nil {
		fmt.Println("Gagal: ", err.Error())
		return
	}

	fmt.Println("Pembayaran Berhasil:", pembayaran.Total, pembayaran.Status)

}
