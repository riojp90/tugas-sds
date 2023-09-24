package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan bilangan (pisahkan dengan spasi): ")
	input, _ := reader.ReadString('\n')

	// Menghapus karakter newline dari input
	input = strings.TrimSpace(input)

	// Memecah input menjadi bilangan-bilangan
	inputSlice := strings.Fields(input)
	numbers := make([]int, len(inputSlice))

	// Mengkonversi string ke int
	for i, str := range inputSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error: Masukkan bilangan bulat yang valid.")
			return
		}
		numbers[i] = num
	}

	// Memeriksa apakah ada elemen dalam slice
	if len(numbers) == 0 {
		fmt.Println("Tidak ada bilangan yang dimasukkan.")
		return
	}

	// Menggunakan fungsi Sort dari package sort untuk mengurutkan bilangan
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	fmt.Println("Bilangan yang telah diurutkan (dari terbesar ke terkecil):", numbers)
}