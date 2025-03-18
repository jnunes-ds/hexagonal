package main

import (
	"database/sql"
	db2 "github.com/jnunes-ds/hexagonal/adapters/db"
	"github.com/jnunes-ds/hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(Db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Example", 19.99)
	productService.Enable(product)
}
