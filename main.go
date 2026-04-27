package main

<<<<<<< HEAD
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
=======
import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Print("bilangan posifi: %d\n", positiveNumber)
	fmt.Print("bilangan negatif: %d\n", negativeNumber)

>>>>>>> ec0880836f30668644dc6556e826781c131e3550
}
