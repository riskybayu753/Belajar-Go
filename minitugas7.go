package main

import (
	"fmt"
)

// Struct Product
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Struct CartItem
type CartItem struct {
	Product  *Product
	Quantity int
	Subtotal float64
}

// Global data
var products = map[int]*Product{}
var productID = 1

func main() {
	for {
		fmt.Println("\n===================================")
		fmt.Println("           MENU UTAMA")
		fmt.Println("===================================")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Lihat Produk")
		fmt.Println("3. Transaksi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addProduct()
		case 2:
			showProducts()
		case 3:
			transaction()
		case 0:
			fmt.Println("Keluar program...")
			return
		default:
			fmt.Println("Menu tidak valid!")
		}
	}
}

// =======================
// TAMBAH PRODUK
// =======================
func addProduct() {
	var name string
	var price float64
	var stock int

	fmt.Print("Nama produk: ")
	fmt.Scanln(&name)

	fmt.Print("Harga: ")
	fmt.Scanln(&price)

	fmt.Print("Stok: ")
	fmt.Scanln(&stock)

	if price <= 0 || stock < 0 {
		fmt.Println("Input tidak valid!")
		return
	}

	products[productID] = &Product{
		ID:    productID,
		Name:  name,
		Price: price,
		Stock: stock,
	}

	fmt.Println("Produk berhasil ditambahkan!")
	productID++
}

// =======================
// LIHAT PRODUK
// =======================
func showProducts() {
	if len(products) == 0 {
		fmt.Println("Belum ada produk!")
		return
	}

	fmt.Println("\n=== DAFTAR PRODUK ===")
	for _, p := range products {
		fmt.Printf("ID: %d | %s | Harga: %.0f | Stok: %d\n",
			p.ID, p.Name, p.Price, p.Stock)
	}
}

// =======================
// TRANSAKSI
// =======================
func transaction() {
	if len(products) == 0 {
		fmt.Println("Tidak ada produk!")
		return
	}

	var cart []CartItem

	for {
		showProducts()

		var id, qty int

		fmt.Print("Masukkan ID produk (0 untuk selesai): ")
		fmt.Scanln(&id)

		if id == 0 {
			break
		}

		product, exists := products[id]
		if !exists {
			fmt.Println("Produk tidak ditemukan!")
			continue
		}

		fmt.Print("Jumlah: ")
		fmt.Scanln(&qty)

		if qty <= 0 || qty > product.Stock {
			fmt.Println("Jumlah tidak valid!")
			continue
		}

		product.Stock -= qty

		subtotal := float64(qty) * product.Price

		cart = append(cart, CartItem{
			Product:  product,
			Quantity: qty,
			Subtotal: subtotal,
		})

		fmt.Println("Ditambahkan ke keranjang")
	}

	if len(cart) == 0 {
		fmt.Println("Tidak ada transaksi")
		return
	}

	// STRUK
	var total float64
	fmt.Println("\n=== STRUK BELANJA ===")
	for _, item := range cart {
		fmt.Printf("%s x%d = %.0f\n",
			item.Product.Name, item.Quantity, item.Subtotal)
		total += item.Subtotal
	}

	fmt.Println("---------------------")
	fmt.Printf("TOTAL: %.0f\n", total)

	// PEMBAYARAN
	var bayar float64
	for {
		fmt.Print("Bayar: ")
		fmt.Scanln(&bayar)

		if bayar < total {
			fmt.Println("Uang kurang!")
		} else {
			break
		}
	}

	if bayar == total {
		fmt.Println("Pembayaran pas")
	} else {
		fmt.Printf("Kembalian: %.0f\n", bayar-total)
	}
}
