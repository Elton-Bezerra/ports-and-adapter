package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Elton-Bezerra/ports-and-adapter/adapters/db"
	"github.com/Elton-Bezerra/ports-and-adapter/app"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		id string, 
		name string, 
		status string, 
		price float);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values 
		("abc", "Product Test", "disabled", 0);`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "abc", product.GetID())
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
	require.Equal(t, 0.0, product.GetPrice())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := app.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Status, productResult.GetStatus())
	require.Equal(t, product.Price, productResult.GetPrice())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.ID, productResult.GetID())
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Status, productResult.GetStatus())
	require.Equal(t, product.Price, productResult.GetPrice())
}
