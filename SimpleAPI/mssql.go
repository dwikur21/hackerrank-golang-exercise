package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

func connectDB() *sql.DB {
	query := url.Values{}
	query.Add("database", "OCBC_EMERGING_DATA")

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword("sa", "deemes"),
		Host:     fmt.Sprintf("%s", "localhost"),
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
