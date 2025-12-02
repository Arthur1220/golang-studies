package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {

	conection := "user=postgres dbname=alura_loja password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)

	if err != nil {
		panic(err.Error())
	}

	return db

}
