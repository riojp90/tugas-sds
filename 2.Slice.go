package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Membuat slice dengan 5 data awal
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice awal:", slice)

	// Meminta input dari pengguna untuk 3 data tambahan
	for i := 0; i < 3; i++ {
		fmt.Print("Masukkan data tambahan ke-", i+1, ": ")
		var input string
		fmt.Scanln(&input)

		// Mengkonversi input menjadi integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Masukkan bilangan bulat yang valid.")
			return
		}

		// Menambahkan data ke dalam slice
		slice = append(slice, num)
	}

	fmt.Println("Slice setelah penambahan data:", slice)
}