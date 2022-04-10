package main

import (
	"github.com/yugendra/TransactionsRoutine/handlers"
	"log"
	"net/http"
)

func main() {
	handler := handlers.NewHandler()
	handler.NewRoutes()

	//Starting the transactionsroutine service
	log.Fatal(http.ListenAndServe(":8080", handler.Router))
}
