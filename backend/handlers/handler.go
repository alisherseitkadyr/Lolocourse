package handlers

import (
	"landing/backend/storage"
	"landing/backend/models"
	"net/http"
	"time"
)

func LeadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	lead := models.Lead{
		FullName:    r.FormValue("fullName"),
		Position:    r.FormValue("position"),
		Company:     r.FormValue("company"),
		Email:       r.FormValue("email"),
		Phone:       r.FormValue("phone"),
		Course_description:     r.FormValue("message"),
		Time:        time.Now().Format(time.RFC3339),
	}

	err := storage.SaveLead(lead)
	if err != nil {
		http.Error(w, "Error saving lead", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))

}

func WithCORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "https://lolocourse.onrender.com")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            return
        }
        next(w, r)
    }
}

