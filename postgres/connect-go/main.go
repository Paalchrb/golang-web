package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type book struct{
	isbn string
	title string
	author string
	price float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://bookie:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database!")

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	books := make([]book, 0)
	for rows.Next() {
		bk := book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price ) //order must match row in DB
		if err != nil {
			panic(err)
		}
		books = append(books, bk)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, bk := range books{
		fmt.Printf("%s %s %s $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}