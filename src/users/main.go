package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Test users service"))
	}).Methods(http.MethodGet)

	modifiedRouter := handlers.LoggingHandler(os.Stdout, router)
	modifiedRouter = handlers.CompressHandler(modifiedRouter)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":3000", modifiedRouter))
}
