package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	var (
		db  *sql.DB
		err error
	)

	// We make connection to the database by mysql giving the appropriate parameters
	db, err = sql.Open("mysql", "root:@/grasfor")
	if err != nil {
		log.Fatal(err)
	}

	/*
		we return the db function of the open connection base, at the moment of creating
		the function a pointer is given to sql.DB to return the functions of this structure
		in the variable db.
	*/
	return db
}
