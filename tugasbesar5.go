// package main

// import (
// 	"fmt"
// )

// // Struct BankAccount
// type BankAccount struct {
// 	AccountNumber string
// 	OwnerName     string
// 	Balance       float64
// 	Transactions  []string
// }

// // Method Deposit (pakai pointer)
// func (b *BankAccount) Deposit(amount float64) {
// 	b.Balance += amount

// 	record := fmt.Sprintf("[DEBIT] +%.0f | Saldo: %.0f", amount, b.Balance)
// 	b.Transactions = append(b.Transactions, record)

// 	fmt.Println("Deposit berhasil:", amount)
// }

// // Method Withdraw (pakai pointer)
// func (b *BankAccount) Withdraw(amount float64) {
// 	if amount > b.Balance {
// 		fmt.Println("Error: Saldo tidak cukup!")
// 		return
// 	}

// 	b.Balance -= amount

// 	record := fmt.Sprintf("[KREDIT] -%.0f | Saldo: %.0f", amount, b.Balance)
// 	b.Transactions = append(b.Transactions, record)

// 	fmt.Println("Withdraw berhasil:", amount)
// }

// // Method GetBalance (tidak perlu pointer)
// func (b BankAccount) GetBalance() {
// 	fmt.Println("Saldo saat ini:", b.Balance)
// }

// // Method PrintStatement
// func (b BankAccount) PrintStatement() {
// 	fmt.Println("\n=== Riwayat Transaksi ===")
// 	for i, t := range b.Transactions {
// 		fmt.Printf("%d. %s\n", i+1, t)
// 	}
// }

// func main() {
// 	// Inisialisasi akun
// 	account := BankAccount{
// 		AccountNumber: "123456789",
// 		OwnerName:     "Ilham",
// 		Balance:       100000,
// 		Transactions:  []string{},
// 	}

// 	// Cek saldo awal
// 	account.GetBalance()

// 	// Deposit
// 	account.Deposit(50000)

// 	// Withdraw berhasil
// 	account.Withdraw(30000)

// 	// Withdraw gagal (saldo tidak cukup)
// 	account.Withdraw(200000)

// 	// Cek saldo akhir
// 	account.GetBalance()

// 	// Tampilkan semua transaksi
// 	account.PrintStatement()
// }
