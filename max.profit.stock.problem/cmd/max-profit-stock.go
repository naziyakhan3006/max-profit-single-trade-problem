package main

import (
	"fmt"
	"log"

	"github.com/julienschmidt/httprouter"
	"net/http"
	"problem.solving/max.profit.stock.problem/handler"
)

func registerEndpoints() *httprouter.Router {
	router := httprouter.New()
	router.GET("/stock/:symbol/days/:days", handler.ProcessSymbol)
	return router
}

func main() {

	router := registerEndpoints()

	fmt.Println(">> listening on localhost:8006")
	log.Println(">> listening on localhost:8006")

	err := http.ListenAndServe("0.0.0.0:8006", router)
	if err != nil {
		log.Fatal(">> failed to start stock service")
	}
}
