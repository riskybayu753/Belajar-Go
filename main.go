package main

import "fmt"

func main() {
	const hargaIndomie = 3500
	var jumlah int = 4
	var uang float64 = 20000
	totalBelanja := hargaIndomie * jumlah
	kembalian := uang - float64(totalBelanja)

	fmt.Printf("Harga Indomie: %d\n", hargaIndomie)
	fmt.Printf("Jumlah beli: %d\n", jumlah)
	fmt.Printf("totalBelanja: %d\n", totalBelanja)
	fmt.Printf("Uang dibayar: %.0f\n", uang)
	fmt.Printf("kembalian: %.0f\n", kembalian)
}
