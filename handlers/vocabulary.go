package handlers

import (
	"encoding/json"
	"github.com/JordenNorton/poly-prompt-backend/db"
	"github.com/JordenNorton/poly-prompt-backend/models"
	"log"
	"net/http"
)

func CreateVocabulary(w http.ResponseWriter, r *http.Request) {
	var vocab models.Vocabulary

	if err := json.NewDecoder(r.Body).Decode(&vocab); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

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
