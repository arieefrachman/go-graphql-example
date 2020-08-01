package models

type Person struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BirthDate string `json:"birth_date"`
}

type PhoneNumber struct {
	ID int `json:"id"`
	PersonID int `json:"person_id"`
	PhoneNumber string `json:"phone_number"`
	Type string `json:"type"`
}