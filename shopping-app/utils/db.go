package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Change these values to your MySQL credentials
	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/yourdatabase" // Replace with your details
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}

