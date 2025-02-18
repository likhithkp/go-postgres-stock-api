package main

import (
	"go-mongo-api/router"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	router.StockRoutes(mux)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
