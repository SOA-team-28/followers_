package main

import (
	"database-example/db"
	"log"
	"net/http"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	http.HandleFunc("/your-route", yourHandlerFunction)

	log.Println("Server is running on port", db.Port)
	log.Fatal(http.ListenAndServe(":"+db.Port, nil))
}

func yourHandlerFunction(w http.ResponseWriter, r *http.Request) {
	// Implementacija tvoje rute
}
