package main

import (
	"fmt"
)

func main() {
	var periode float64

	// Meminta input dari pengguna untuk periode dalam detik
	fmt.Print("Masukkan periode dalam detik: ")
	fmt.Scan(&periode)

	// Menghitung frekuensi
	frekuensi := 1 / periode

	// Menampilkan hasil
	fmt.Printf("Frekuensi: %.2f Hz\n", frekuensi)
}