package service

import (
	"go-mongo-api/db"
	"go-mongo-api/model"
	"log"
)

func InsertStock(stock *model.Stock) int64 {
	defer db.CreateConnection().Close()
	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`
	var id int64
	err := db.CreateConnection().QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatal("Unable to execute the query", "%v", err)
	}
	return id
}

func GetStock(id int64) (model.Stock, error) {
	defer db.CreateConnection().Close()
	var stock model.Stock
	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`
	row := db.CreateConnection().QueryRow(sqlStatement, id)

	err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
	return stock, err
}

func GetAllStock() ([]model.Stock, error) {
	defer db.CreateConnection().Close()

	var stocks []model.Stock
	sqlStatement := `SELECT * FROM stocks;`
	rows, err := db.CreateConnection().Query(sqlStatement)

	if err != nil {
		log.Fatal("Unable to execute the query")
	}

	defer rows.Close()

	for rows.Next() {
		var stock model.Stock
		rows.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
		stocks = append(stocks, stock)
	}

	return stocks, err
}

func UpdateStock(id int64, stock *model.Stock) {
	defer db.CreateConnection().Close()
	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`
	res, err := db.CreateConnection().Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatal("Unable to execute the query")
	}

	res.RowsAffected()

	if err != nil {
		log.Fatal("Unable to check rows effected")
	}
}

func DeleteStock(id int64) {
	defer db.CreateConnection().Close()
	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`
	res, err := db.CreateConnection().Exec(sqlStatement, id)

	if err != nil {
		log.Fatal("Unable to execute the query")
	}

	res.RowsAffected()

	if err != nil {
		log.Fatal("Unable to check rows effected")
	}
}
