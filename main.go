package main

import (
	"fmt"
)

func tambah(a, b float64) float64 {
	return a + b
}
func kurang(a, b float64) float64 {
	return a - b
}
func kali(a, b float64) float64 {
	return a * b
}
func bagi(a, b float64) float64 {
	if b == 0 {
		fmt.Println("❌ Tidak bisa dibagi dengan 0")
		return 0
	}
	return a / b
}

func main() {
	var a, b float64
	var operator string

	fmt.Print("masukan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Printf("masukan operator \n1 +\n2 -\n3 *\n4 /\npilih sesuai angka atau operator diatas :")
	fmt.Scanln(&operator)
	fmt.Print("masukan angka kedua: ")
	fmt.Scanln(&b)
	fmt.Println("============================")

	var hasil float64

	switch operator {
	case "1", "+":
		operator = "+"
		hasil = tambah(a, b)
	case "2", "-":
		operator = "-"
		hasil = kurang(a, b)
	case "3", "*":
		operator = "*"
		hasil = kali(a, b)
	case "4", "/":
		operator = "/"
		hasil = bagi(a, b)
	default:
		fmt.Println("❌ Operator tidak dikenal")
		return
	}
	fmt.Printf("Hasil Perhitungan yaitu %.2f %s %.2f = %.2f\n", a, operator, b, hasil)
}
