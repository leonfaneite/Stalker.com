package dbpostgrest

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "postgresql"
	port     = 5432
	user     = "usuario"
	password = "usuario"
	dbname   = "customer"
)

func GetConnect() *sql.DB {

	connect := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db

}
