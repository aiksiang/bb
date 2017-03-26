package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //
)

// Save .
func Save() {
	fmt.Println("Saving...")
	db, err := sql.Open("mysql", "root:oK3o=ktGJSb4@tcp(127.0.0.1:3306)/bb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select * from crawl_raw")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int
	var blogShop string
	var itemData string

	for rows.Next() {
		if err := rows.Scan(&id, &blogShop, &itemData); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, blogShop, itemData)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// Create .
func Create() {

}
