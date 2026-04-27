package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_sample")
	if err != nil {
		return nil, err
	}

	fmt.Println("Berhasil terhubung ke database!")

	return db, nil
}

func main() {
	db, err := connect()
	if err != nil {
		fmt.Println("Gagal terhubung ke database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Selesai.....")
}
