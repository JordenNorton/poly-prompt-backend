package main

import (
	"fmt"
	"github.com/JordenNorton/poly-prompt-backend/config"
	"github.com/JordenNorton/poly-prompt-backend/db"
	"log"
	"net/http"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	// Initialise the database connection
	if err := db.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
