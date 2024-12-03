package models

type PatientInfo struct {
	ID          int    `json:"id"`
	PatientID   string `json:"patient_id"`
	Image       []byte `json:"image"`
	Name        string `json:"name_"`
	PhoneNumber string `json:"phone_numbers"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	CardID      string `json:"card_id"`
	Religion    string `json:"religion"`
	Sex         string `json:"sex"`
	Blood       string `json:"blood"`
	DateOfBirth string `json:"date_of_birth"`
	Age         int    `json:"age"`
}
