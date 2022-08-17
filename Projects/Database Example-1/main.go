package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func main() {

}

func select_all_rows(db *sql.DB) {
	response, err := db.Query("SELECT * FROM  cities")
	defer db.Close()
	if err != nil {
		fmt.Println("HATA")
	}

	query := "UPDATE users SET name=$1, password=$2"
	a, err := db.Exec(query, "mustafa", "123")
	fmt.Println(a)

	for response.Next() {
		var city City
		err := response.Scan(&city.Id, &city.Name, &city.Population)
		if err != nil {
			fmt.Println("HATA")
		}
		fmt.Println(city)
	}

}

func conn() *sql.DB {
	db, err := sql.Open("mysql", "root:.......")

	if err != nil {
		fmt.Println("HATA")
	}
	return db
}

type City struct {
	Id         int
	Name       string
	Population string
}
