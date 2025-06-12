package main

import (
	"log"
	"net/http"

	"github.com/cristina-sirbu/loan-buddy/api"
)

func main() {
	http.HandleFunc("/offers", api.OffersHandler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
