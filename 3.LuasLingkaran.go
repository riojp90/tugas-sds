package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi variabel untuk menyimpan jari-jari lingkaran
	var jariJari float64

	// Meminta input dari pengguna untuk jari-jari lingkaran
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jariJari)

	// Menghitung luas lingkaran
	luas := math.Pi * jariJari * jariJari

	// Menampilkan hasil
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, luas)
}