package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abraaoduarte/ekklesia-backend-golang/internal/database"
	"github.com/abraaoduarte/ekklesia-backend-golang/pkg/server"
)

func main() {
	err := database.Init()

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	router := server.NewRouter()

	server.RegisterHandlers(router)


	address := fmt.Sprintf(":%v", os.Getenv("PORT"))

	log.Printf("Server is running on %s...\n", address)

	err = http.ListenAndServe(address, router)

	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}