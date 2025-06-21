package storage

import (
	"landing/backend/models"
)

func SaveLead(lead models.Lead) error {
	_, err := DB.Exec(`INSERT INTO leads (name, phone, email, time) VALUES (?, ?, ?, ?)`,
		lead.Name, lead.Phone, lead.Email, lead.Time)
	return err
}
