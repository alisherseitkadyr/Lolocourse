package handlers

import (
	"encoding/json"
	"landing/backend/models"
	"landing/backend/storage"
	"net/http"
	"time"
)

func HandleLead(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var lead models.Lead
	if err := json.NewDecoder(r.Body).Decode(&lead); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	lead.Time = time.Now().Format(time.RFC3339)

	if err := storage.SaveLead(lead); err != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Заявка сохранена",
	})
}
