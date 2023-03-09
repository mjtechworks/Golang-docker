package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "test_erp"
)

func Init() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {

		return nil
	}

	fmt.Println("Successfully connected!")

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil
	}

	// row, err := db.Query("SELECT * FROM users;")

	// if err != nil {
	// 	fmt.Println(nil)
	// } else {

	// 	fmt.Println("<<<<<<<<<<<<", row)
	// }

	return db

}
