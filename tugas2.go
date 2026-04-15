package main

import (
	"fmt"
	"strings"
)

func main() {
	cart := []string{}

	products := map[string]int{
		"indomie": 3000,
		"susu":    5000,
		"roti":    4000,
	}

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Lihat Produk")
		fmt.Println("2. Tambah ke belanja")
		fmt.Println("3. Lihat belanja")
		fmt.Println("4. Hitung Total belanja")
		fmt.Println("5. Exit")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {

		// STEP 1
		case 1:
			fmt.Println("\nDaftar Produk:")
			for name, price := range products {
				fmt.Println(name, "-", price)
			}

		// STEP 2
		case 2:
			var input string
			fmt.Print("Masukkan produk: ")
			fmt.Scanln(&input)

			cart = append(cart, input)
			fmt.Println("Produk ditambahkan!")

		// STEP 3
		case 3:
			fmt.Println("\nIsi Cart:")
			if len(cart) == 0 {
				fmt.Println("Cart masih kosong")
			} else {
				for i, item := range cart {
					fmt.Println(i+1, ".", item)
				}
			}

		// STEP 4
		case 4:
			total := 0

			for _, item := range cart {
				key := strings.ToLower(item)

				price, ok := products[key]
				if ok {
					total += price
				} else {
					fmt.Println("Produk tidak ditemukan:", item)
				}
			}

			fmt.Printf("Total: Rp%d\n", total)

		// STEP 5 (EXIT)
		case 5:
			fmt.Println("Terima kasih sudah berkunjung")
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
