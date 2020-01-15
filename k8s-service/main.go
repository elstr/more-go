package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/elstr/more-go/k8s-service/handlers"
)

func main() {
	log.Print("Starting the service...")
	router := handlers.Router()	
	log.Print("The service is ready to listen and serve.")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
