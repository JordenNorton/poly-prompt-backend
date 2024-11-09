package models

type Vocabulary struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Word        string `json:"word"`
	Translation string `json:"translation"`
	Difficulty  int    `json:"difficulty"`
}
