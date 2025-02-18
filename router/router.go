package router

import (
	"go-mongo-api/handler"
	"net/http"
)

func StockRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/allStocks", handler.GetAllStocks)
	mux.HandleFunc("/getstock", handler.UpdateStock)
	mux.HandleFunc("/createStock", handler.CreateStock)
	mux.HandleFunc("/updateStock", handler.UpdateStock)
	mux.HandleFunc("/deleteStock", handler.DeleteStock)
}
