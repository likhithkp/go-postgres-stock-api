package handler

import (
	"encoding/json"
	"go-mongo-api/db"
	"go-mongo-api/model"
	"go-mongo-api/service"
	"log"
	"net/http"
	"strconv"
)

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var newStock *model.Stock
	err := json.NewDecoder(r.Body).Decode(&newStock)

	if err != nil {
		log.Fatal("Unable to decode", "%v", err)
	}

	insertId := service.InsertStock(newStock)

	res := db.Response{
		ID:      insertId,
		Message: "Stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stock, err := service.GetAllStock()

	if err != nil {
		log.Fatal("Unable to get all stocks", "%v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	stockId, _ := strconv.Atoi(id)
	stock, err := service.GetStock(int64(stockId))

	if err != nil {
		log.Fatal("Unable to find the stock", "%v", err)
	}

	json.NewEncoder(w).Encode(stock)

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	stockId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
		return
	}

	service.DeleteStock(int64(stockId))

	res := db.Response{
		ID:      int64(stockId),
		Message: "Deleted",
	}

	json.NewEncoder(w).Encode(res)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	stockId, err := strconv.Atoi(id)

	var updatedStock model.Stock
	err = json.NewDecoder(r.Body).Decode(&updatedStock)

	if err != nil {
		log.Fatal("Unable to decode the stock", "%v", err)
	}

	service.UpdateStock(int64(stockId), &updatedStock)

	res := db.Response{
		ID:      int64(stockId),
		Message: "Updated",
	}

	json.NewEncoder(w).Encode(res)
}
