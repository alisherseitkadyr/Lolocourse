package storage

import (
	"landing/backend/models"
)

func SaveLead(lead models.Lead) error {
	_, err := DB.Exec(`INSERT INTO leads (fullname, position, company, email, phone, course_description, time) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		lead.FullName, lead.Position, lead.Company, lead.Email, lead.Phone, lead.Course_description, lead.Time)
	return err
}
