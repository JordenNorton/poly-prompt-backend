package main

import (
	"github.com/JordenNorton/poly-prompt-backend/config"
	"github.com/JordenNorton/poly-prompt-backend/db"
	"github.com/JordenNorton/poly-prompt-backend/handlers"
	"github.com/gorilla/mux"
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

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")
	router.HandleFunc("/vocabularies", handlers.GetAllVocabulary).Methods("GET")
	router.HandleFunc("/vocabulary", handlers.CreateVocabulary).Methods("POST")
	router.HandleFunc("/vocabulary/{id}", handlers.UpdateVocabulary).Methods("PUT")
	router.HandleFunc("/vocabulary/{id}", handlers.DeleteVocabulary).Methods("DELETE")

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
