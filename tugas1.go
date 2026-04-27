package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id     int
	Name   string
	Email  string
	Status int
	Phone  string
}

// ================= DATABASE =================

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/latihan_go?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Berhasil terhubung ke database!")
	return db, nil
}

// ================= CRUD =================

func add(db *gorm.DB, data User) error {
	tx := db.Exec("INSERT INTO users (name, email, status, phone) VALUES (?, ?, ?, ?)",
		data.Name, data.Email, data.Status, data.Phone)

	return tx.Error
}

func getAll(db *gorm.DB) ([]User, error) {
	var users []User
	tx := db.Raw("SELECT * FROM users").Scan(&users)
	return users, tx.Error
}

func getById(db *gorm.DB, id int) (User, error) {
	var user User
	tx := db.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)
	return user, tx.Error
}

func update(db *gorm.DB, id int, data User) error {
	tx := db.Exec("UPDATE users SET name=?, email=?, status=?, phone=? WHERE id=?",
		data.Name, data.Email, data.Status, data.Phone, id)

	return tx.Error
}

func deleteUser(db *gorm.DB, id int) error {
	tx := db.Exec("DELETE FROM users WHERE id = ?", id)

	if tx.RowsAffected == 0 {
		return fmt.Errorf("data tidak ditemukan")
	}

	return tx.Error
}

// ================= MAIN MENU =================

func main() {
	db, err := connectGorm()
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\npengelolaan user")
		fmt.Println("1. lihat semua user")
		fmt.Println("2. tambah user")
		fmt.Println("3. lihat detail user")
		fmt.Println("4. update user")
		fmt.Println("5. hapus user")
		fmt.Println("0. keluar")

		fmt.Print("pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		case "1":
			users, err := getAll(db)
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, u := range users {
				fmt.Printf("ID: %d, Name: %s, Email: %s, Status: %d, Phone: %s\n",
					u.Id, u.Name, u.Email, u.Status, u.Phone)
			}

		case "2":
			var user User

			fmt.Print("masukkan nama: ")
			user.Name, _ = reader.ReadString('\n')
			user.Name = strings.TrimSpace(user.Name)

			fmt.Print("masukkan email: ")
			user.Email, _ = reader.ReadString('\n')
			user.Email = strings.TrimSpace(user.Email)

			fmt.Print("masukkan phone: ")
			user.Phone, _ = reader.ReadString('\n')
			user.Phone = strings.TrimSpace(user.Phone)

			user.Status = 1

			err := add(db, user)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("data berhasil ditambahkan")
			}

		case "3":
			fmt.Print("masukkan id: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			user, err := getById(db, id)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("ID: %d, Name: %s, Email: %s, Status: %d, Phone: %s\n",
				user.Id, user.Name, user.Email, user.Status, user.Phone)

		case "4":
			fmt.Print("masukkan id: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			var user User

			fmt.Print("nama baru: ")
			user.Name, _ = reader.ReadString('\n')
			user.Name = strings.TrimSpace(user.Name)

			fmt.Print("email baru: ")
			user.Email, _ = reader.ReadString('\n')
			user.Email = strings.TrimSpace(user.Email)

			fmt.Print("phone baru: ")
			user.Phone, _ = reader.ReadString('\n')
			user.Phone = strings.TrimSpace(user.Phone)

			user.Status = 1

			err := update(db, id, user)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("data berhasil diupdate")
			}

		case "5":
			fmt.Print("masukkan id: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			err := deleteUser(db, id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("data berhasil dihapus")
			}

		case "0":
			fmt.Println("keluar...")
			return

		default:
			fmt.Println("pilihan tidak valid")
		}
	}
}
