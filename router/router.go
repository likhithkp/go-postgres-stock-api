package router

import (
	"go-mongo-api/handler"
	"net/http"
)

func StockRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /stocks", handler.GetAllStocks)
	mux.HandleFunc("GET /stocks/{id}", handler.GetStock)
	mux.HandleFunc("/stocks", handler.CreateStock)
	mux.HandleFunc("/stocks/{id}", handler.UpdateStock)
	mux.HandleFunc("DELETE /stocks/{id}", handler.DeleteStock)
}
