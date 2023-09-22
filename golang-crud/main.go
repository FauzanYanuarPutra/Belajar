package main

import (
	"golang-crud/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	
	log.Println("Server running on port 8080")
	http.ListenAndServe(":3000", nil)
}