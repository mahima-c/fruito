package model

import "time"

type Patient struct {
	ID               int              `json:"id"`
	PatientID        string           `json:"patient_id"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	DateOfBirth      *time.Time       `json:"date_of_birth"`
	Gender           string           `json:"gender"`
	Address          string           `json:"address"`
	City             string           `json:"city"`
	State            string           `json:"state"`
	ZipCode          string           `json:"zip_code"`
	Country          string           `json:"country"`
	PhoneNumber      string           `json:"phone_number"`
	Email            string           `json:"email"`
	EmergencyContact EmergencyContact `json:"emergency_contact"`
	Allergies        []Allergie       `json:"allergies,omitempty"`
	Notes            string           `json:"notes,omitempty"`
	AdmittedAt       *time.Time       `json:"admitted_at,omitempty"`
	DischargedAt     *time.Time       `json:"discharged_at,omitempty"`
	WardID           int              `json:"ward_id,omitempty"`
	RoomID           int              `json:"room_id,omitempty"`
	BedID            int              `json:"bed_id,omitempty"`
}

type EmergencyContact struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email,omitempty"`
}

type Allergie struct {
	Allergen  string `json:"allergen"`
	Severity  string `json:"severity"`  // e.g., Mild, Moderate, Severe
	Reactions string `json:"reactions"` // e.g., Rash, Swelling, Difficulty breathing
}
