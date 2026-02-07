package model

import "time"

type MedicalHistory struct {
	ID            int             `json:"id"`
	PatientID     int             `json:"patiend_id"` // patient table foreign key
	Conditions    []Condition     `json:"conditions"`
	Surgeries     []Surgery       `json:"surgeries"`
	Medications   []Medication    `json:"medications"`
	Immunizations []Immunization  `json:"immunizations"`
	FamilyHistory []FamilyHistory `json:"family_history"`
}

type Condition struct {
	Name                 string       `json:"name"`
	DiagnosisDate        *time.Time   `json:"diagnosis_date,omitempty"`
	DiagnosisReportLinks []ReportLink `json:"diagnosis_report_links,omitempty"`
	Notes                string       `json:"notes,omitempty"`
}

type Surgery struct {
	Name               string       `json:"name"`
	Date               time.Time    `json:"date,omitempty"`
	SurgeryReportLinks []ReportLink `json:"surgery_report_links,omitempty"`
	Notes              string       `json:"notes,omitempty"`
}

type Medication struct {
	Name      string    `json:"name"`
	Dosage    string    `json:"dosage,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Notes     string    `json:"notes,omitempty"`
}

type Immunization struct {
	Name string    `json:"name"`
	Date time.Time `json:"date,omitempty"`
}

type FamilyHistory struct {
	Relative  string `json:"relative"`
	Condition string `json:"condition"`
}

type ReportLink struct {
	ReportId   int    `json:"report_id"`
	ReportType string `json:"report_type"`
	ReportURL  string `json:"report_url"`
}
