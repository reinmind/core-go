package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	PartnerId   string `json:"partner_id"`
	CompanyName string `json:"company_name"`
}

func main() {
	LogSettings()
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "hyperchain:B1ockch@inhyperchain@tcp(127.0.0.1:3307)/nbdc_open")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	results, err := db.Query("select partner_id,company_name from company")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.PartnerId, &tag.CompanyName)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.CompanyName)
		log.Printf(tag.PartnerId)
	}
}

func LogSettings() {
	log.SetPrefix("sql-test:")
	log.SetFlags(log.LstdFlags)
}
