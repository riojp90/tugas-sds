package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk menyimpan panjang alas dan tinggi
	var alas, tinggi float64

	// Meminta input dari pengguna untuk panjang alas
	fmt.Print("Masukkan panjang alas segitiga: ")
	fmt.Scan(&alas)

	// Meminta input dari pengguna untuk tinggi segitiga
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	// Menghitung luas segitiga
	luas := 0.5 * alas * tinggi

	// Menampilkan hasil
	fmt.Printf("Luas segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f\n", alas, tinggi, luas)
}
