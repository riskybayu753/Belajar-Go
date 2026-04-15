// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// // Struct TicketSystem
// type TicketSystem struct {
// 	TotalTickets     int
// 	RemainingTickets int
// 	mu               sync.Mutex
// }

// // Method BookTicket
// func (t *TicketSystem) BookTicket(buyerName string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	// Simulasi delay random (100 - 300 ms)
// 	time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)

// 	t.mu.Lock()
// 	defer t.mu.Unlock()

// 	if t.RemainingTickets > 0 {
// 		t.RemainingTickets--
// 		fmt.Printf("✅ %s berhasil pesan tiket | Sisa: %d\n", buyerName, t.RemainingTickets)
// 	} else {
// 		fmt.Printf("❌ %s gagal, tiket habis\n", buyerName)
// 	}
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	// Inisialisasi sistem tiket
// 	ticketSystem := TicketSystem{
// 		TotalTickets:     10,
// 		RemainingTickets: 10,
// 	}

// 	// Daftar 15 pembeli
// 	buyers := []string{
// 		"Ilham", "Budi", "Siti", "Andi", "Rina",
// 		"Dewi", "Agus", "Putra", "Lina", "Doni",
// 		"Rudi", "Tono", "Joko", "Ayu", "Nina",
// 	}

// 	var wg sync.WaitGroup

// 	// Jalankan goroutine untuk tiap pembeli
// 	for _, buyer := range buyers {
// 		wg.Add(1)
// 		go ticketSystem.BookTicket(buyer, &wg)
// 	}

// 	// Tunggu semua selesai
// 	wg.Wait()

// 	// Ringkasan
// 	sold := ticketSystem.TotalTickets - ticketSystem.RemainingTickets

// 	fmt.Println("\n=== Ringkasan ===")
// 	fmt.Println("Total tiket :", ticketSystem.TotalTickets)
// 	fmt.Println("Terjual     :", sold)
// 	fmt.Println("Sisa tiket  :", ticketSystem.RemainingTickets)
// 	fmt.Println("Gagal beli  :", len(buyers)-sold)
// }
