package main

import "fmt"

func Latihan1() {
	var namaProject string = "GO App"
	var tahunRilis int = 2022
	var isSaas bool = true

	if tahunRilis < 2024 {
		fmt.Println("Project Lama", namaProject)
	} else {
		fmt.Println("Project Baru", namaProject)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "(Genap)")
		} else {
			fmt.Println(i)
		}
	}

	status := cekStatus(isSaas)

	fmt.Println(status)
}

func cekStatus(appStatus bool) string {
	if appStatus {
		return "Aplikasi SaaS"
	} else {
		return "Aplikasi Project"
	}
}
