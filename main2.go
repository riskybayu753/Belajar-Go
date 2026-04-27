// package main

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	Id     int
// 	Name   string
// 	Email  string
// 	Status int
// 	Phone  string
// }

// func connectGorm() (*gorm.DB, error) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Berhasil terhubung ke database!")

// 	return db, nil
// }

// func main() {
// 	db, err := connectGorm()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// Menambahkan data baru
// 	// newUser := User{
// 	// 	Name:   "John Daa",
// 	// 	Email:  "jandaa@example.com",
// 	// 	Status: 0,
// 	// 	Phone:  "08123456709",
// 	// }

// 	// err = add(db, newUser)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// 	return
// 	// }

// 	// Mendapatkan semua data
// 	users, err := getAll(db)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	for _, user := range users {
// 		fmt.Printf("ID: %d, Name: %s, Email: %s, Status: %d, Phone: %s\n",
// 			user.Id, user.Name, user.Email, user.Status, user.Phone)
// 	}

// 	// Update data
// 	updateData := User{
// 		Name:   "John Smith",
// 		Email:  "smith@example.com",
// 		Status: 1,
// 		Phone:  "08125335213",
// 	}

// 	err = update(db, 6, updateData)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// Mendapatkan data berdasarkan ID
// 	user, err := getById(db, 6)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Printf("ID: %d, Name: %s, Email: %s, Status: %d, Phone: %s\n",
// 		user.Id, user.Name, user.Email, user.Status, user.Phone)
// }

// func add(db *gorm.DB, data User) error {
// 	tx := db.Exec("insert into users (name, email, status, phone) VALUES (?, ?, ?, ?)",
// 		data.Name, data.Email, data.Status, data.Phone)
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	fmt.Println("Data berhasil disimpan")
// 	return nil
// }

// func update(db *gorm.DB, id int, data User) error {
// 	tx := db.Exec("UPDATE users SET name=?, email=?, status=?, phone=? WHERE id=?",
// 		data.Name, data.Email, data.Status, data.Phone, id)
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	fmt.Println("Data berhasil diperbarui")
// 	return nil
// }

// func getAll(db *gorm.DB) ([]User, error) {
// 	var users []User
// 	tx := db.Raw("select * from users").Scan(&users)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	fmt.Println("Data berhasil diambil")
// 	return users, nil
// }

// func getById(db *gorm.DB, id int) (User, error) {
// 	var user User
// 	tx := db.Raw("select * from users where id = ?", id).Scan(&user)
// 	if tx.Error != nil {
// 		return User{}, tx.Error
// 	}

// 	fmt.Printf("Data %d berhasil diambil\n", id)
// 	return user, nil
// }
