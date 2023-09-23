package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk menyimpan panjang sisi
	var sisi int

	// Meminta input dari pengguna
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&sisi)

	// Menghitung luas persegi
	luas := sisi * sisi

	// Menampilkan hasil
	fmt.Printf("Luas persegi dengan sisi %d adalah %d\n", sisi, luas)
}