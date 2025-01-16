package db

import (
	"database/sql"
	"fmt"
)
const (
	host = "localhost"
	port 5432
	user = "postgres"
	password = "1234"
	dbname = "postgres"
)

func ConnectDB() {
	// Connect to the "bank" database.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s  password=%s sslmode=disable"+
	host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to " + dbname)