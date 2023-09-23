package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi variabel
	var massa, tinggi, kecepatan float64

	// Meminta input dari pengguna untuk massa (kg) dan tinggi (m)
	fmt.Print("Masukkan massa objek (kg): ")
	fmt.Scan(&massa)
	fmt.Print("Masukkan tinggi objek (m): ")
	fmt.Scan(&tinggi)

	// Meminta input dari pengguna untuk kecepatan (m/s)
	fmt.Print("Masukkan kecepatan objek (m/s): ")
	fmt.Scan(&kecepatan)

	// Menghitung energi potensial gravitasi
	g := 9.81 // Percepatan gravitasi bumi (m/s^2)
	energiPotensial := massa * g * tinggi

	// Menghitung energi kinetik
	energiKinetik := 0.5 * massa * math.Pow(kecepatan, 2)

	// Menampilkan hasil
	fmt.Printf("Energi Potensial Gravitasi: %.2f joule\n", energiPotensial)
	fmt.Printf("Energi Kinetik: %.2f joule\n", energiKinetik)
}