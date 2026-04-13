package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Golang Basic Learning ===")
	fmt.Println("Pilih latihan:")
	fmt.Println("1. Latihan 1 - Variabel, Kondisi, Perulangan, Fungsi")
	fmt.Println("2. Latihan 2 - Struct & Pointer")
	fmt.Println("3. Latihan 3 - Struct, Error Handling, Multi-return")
	fmt.Println("4. Latihan 4 - Slice of Struct & Range")
	fmt.Print("\nMasukkan pilihan (1-4): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println()

	switch input {
	case "1":
		Latihan1()
	case "2":
		Latihan2()
	case "3":
		Latihan3()
	case "4":
		Latihan4()
	default:
		fmt.Println("Pilihan tidak valid. Masukkan angka 1-4.")
	}
}
