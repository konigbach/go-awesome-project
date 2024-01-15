package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"os"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "db"
)

func insertRecord(db *sql.DB, record []string) error {
	// Assuming your table has three columns: col1, col2, col3
	query := "INSERT INTO names(name, age) VALUES($1, $2)"
	_, err := db.Exec(query, record[0], record[1])
	return err
}

func main() {
	file, err := os.Open("csv/subjects.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Connect to the PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	// Read records and insert into the PostgreSQL database
	for {
		record, err1 := reader.Read()
		if err1 != nil {
			// Check for end of file
			if err1 == io.EOF {
				break
			}
			fmt.Println("Error reading CSV:", err1)
			return
		}

		// Insert the record into the database
		err = insertRecord(db, record)
		if err != nil {
			fmt.Println("Error inserting record into the database:", err)
			return
		}
	}
}
