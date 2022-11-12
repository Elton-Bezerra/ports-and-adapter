package main

import (
	"database/sql"

	db2 "github.com/Elton-Bezerra/ports-and-adapter/adapters/db"
	"github.com/Elton-Bezerra/ports-and-adapter/app"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := app.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 300.0)

	productService.Enable(product)
}
