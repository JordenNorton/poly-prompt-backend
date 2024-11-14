package handlers

import (
	"encoding/json"
	"github.com/JordenNorton/poly-prompt-backend/db"
	"github.com/JordenNorton/poly-prompt-backend/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func CreateVocabulary(w http.ResponseWriter, r *http.Request) {

	validTypes := map[string]bool{
		"noun":      true,
		"verb":      true,
		"adjective": true,
		"adverb":    true,
	}

	var vocab models.Vocabulary

	// Decode the request body into the vocab struct
	if err := json.NewDecoder(r.Body).Decode(&vocab); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Input validation
	if vocab.Type == "" || vocab.Word == "" || vocab.Translation == "" {
		http.Error(w, "Missing required fields: type, word and translation are mandatory", http.StatusBadRequest)
		return
	}
	if vocab.Difficulty < 1 || vocab.Difficulty > 5 {
		http.Error(w, "Invalid difficulty: must be between 1 and 5", http.StatusBadRequest)
		return
	}
	if !validTypes[vocab.Type] {
		http.Error(w, "Invalid type: must be one of 'noun', 'verb', 'adjective', or 'adverb'", http.StatusBadRequest)
		return
	}

	// Insert into the database
	query := `INSERT INTO vocabulary (type, word, translation, difficulty) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.DB.QueryRow(query, vocab.Type, vocab.Word, vocab.Translation, vocab.Difficulty).Scan(&vocab.ID)
	if err != nil {
		log.Println("Error inserting vocabulary")
		http.Error(w, "Failed to create entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vocab)
}

func GetAllVocabulary(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(`SELECT id, type, word, translation, difficulty FROM vocabulary`)
	if err != nil {
		log.Println("Error fetching vocabulary")
		http.Error(w, "Failed to retrieve entries", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var vocabList []models.Vocabulary
	for rows.Next() {
		var vocab models.Vocabulary
		if err := rows.Scan(&vocab.ID, &vocab.Type, &vocab.Word, &vocab.Translation, &vocab.Difficulty); err != nil {
			log.Println("Error scanning row", err)
			continue
		}
		vocabList = append(vocabList, vocab)
	}

	json.NewEncoder(w).Encode(vocabList)

}

func UpdateVocabulary(w http.ResponseWriter, r *http.Request) {

	validTypes := map[string]bool{
		"noun":      true,
		"verb":      true,
		"adjective": true,
		"adverb":    true,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var vocab models.Vocabulary
	if err := json.NewDecoder(r.Body).Decode(&vocab); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate fields
	if vocab.Type == "" || vocab.Word == "" || vocab.Translation == "" {
		http.Error(w, "Missing required fields: type, word, and translation are mandatory", http.StatusBadRequest)
		return
	}
	if vocab.Difficulty < 1 || vocab.Difficulty > 5 {
		http.Error(w, "Invalid difficulty: must be between 1 and 5", http.StatusBadRequest)
		return
	}
	if !validTypes[vocab.Type] {
		http.Error(w, "Invalid type: must be one of 'noun', 'verb', 'adjective', or 'adverb'", http.StatusBadRequest)
		return
	}

	// Update the database
	query := `UPDATE vocabulary SET type = $1, word = $2, translation = $3, difficulty = $4 WHERE id = $5`
	_, err = db.DB.Exec(query, vocab.Type, vocab.Word, vocab.Translation, vocab.Difficulty, id)
	if err != nil {
		log.Printf("Error updating vocabulary: %v\n", err)
		http.Error(w, "Failed to update entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vocabulary updated successfully"})
}

func DeleteVocabulary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM vocabulary WHERE id = $1`
	_, err = db.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting vocabulary: %v\n", err)
		http.Error(w, "Failed to delete entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vocabulary deleted successfully"})
}
