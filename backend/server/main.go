package main

import (
	"log"
	"net/http"

	"github.com/gerdalukosiute/fountains/internal/api"
	"github.com/gerdalukosiute/fountains/internal/database"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.InitDB("fountains.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	api.SetupRoutes(router, db)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
