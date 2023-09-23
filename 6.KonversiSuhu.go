package main

import (
	"fmt"
)

func main() {
	var suhu float64
	var unitAsal, unitTujuan string

	fmt.Print("Masukkan suhu: ")
	fmt.Scan(&suhu)
	fmt.Print("Masukkan unit suhu (C/F/K): ")
	fmt.Scan(&unitAsal)
	fmt.Print("Masukkan unit tujuan (C/F/K): ")
	fmt.Scan(&unitTujuan)

	switch unitAsal {
	case "C":
		switch unitTujuan {
		case "C":
			fmt.Printf("%.2f °C sama dengan %.2f °C\n", suhu, suhu)
		case "F":
			fahrenheit := (suhu * 9/5) + 32
			fmt.Printf("%.2f °C sama dengan %.2f °F\n", suhu, fahrenheit)
		case "K":
			kelvin := suhu + 273.15
			fmt.Printf("%.2f °C sama dengan %.2f K\n", suhu, kelvin)
		default:
			fmt.Println("Unit suhu tujuan tidak valid.")
		}
	case "F":
		switch unitTujuan {
		case "C":
			celsius := (suhu - 32) * 5/9
			fmt.Printf("%.2f °F sama dengan %.2f °C\n", suhu, celsius)
		case "F":
			fmt.Printf("%.2f °F sama dengan %.2f °F\n", suhu, suhu)
		case "K":
			kelvin := (suhu - 32) * 5/9 + 273.15
			fmt.Printf("%.2f °F sama dengan %.2f K\n", suhu, kelvin)
		default:
			fmt.Println("Unit suhu tujuan tidak valid.")
		}
	case "K":
		switch unitTujuan {
		case "C":
			celsius := suhu - 273.15
			fmt.Printf("%.2f K sama dengan %.2f °C\n", suhu, celsius)
		case "F":
			fahrenheit := (suhu-273.15)*9/5 + 32
			fmt.Printf("%.2f K sama dengan %.2f °F\n", suhu, fahrenheit)
		case "K":
			fmt.Printf("%.2f K sama dengan %.2f K\n", suhu, suhu)
		default:
			fmt.Println("Unit suhu tujuan tidak valid.")
		}
	default:
		fmt.Println("Unit suhu asal tidak valid.")
	}
}