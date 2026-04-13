// Menu interaktif untuk memilih dan menjalankan latihan Go.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Membuat reader untuk membaca input dari user
	reader := bufio.NewReader(os.Stdin)

	// Tampilkan menu pilihan latihan
	fmt.Println("=== Golang Basic Learning ===")
	fmt.Println("Pilih latihan:")
	fmt.Println("1. Latihan 1 - Variabel, Kondisi, Perulangan, Fungsi")
	fmt.Println("2. Latihan 2 - Struct & Pointer")
	fmt.Println("3. Latihan 3 - Struct, Error Handling, Multi-return")
	fmt.Println("4. Latihan 4 - Slice of Struct & Range")
	fmt.Print("\nMasukkan pilihan (1-4): ")

	// Baca dan bersihkan input user
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println()

	// Jalankan latihan sesuai pilihan user
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
