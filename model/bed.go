package model

type Bed struct {
	ID          int      `json:"id"`
	WardID      int      `json:"ward_id"` // Foreign key to Ward
	BedNumber   string   `json:"bed_number"`
	BedLocation Location `json:"location"`
	Type        string   `json:"type"`                   // e.g., Standard, ICU, Pediatric, Bariatric
	Status      string   `json:"status"`                 // e.g., Available, Occupied, Maintenance, Reserved
	OccupiedBy  int      `json:"occupied_by,omitempty"`  // PatientID (foreign key to Patient), omit if not occupied
	LastCleaned string   `json:"last_cleaned,omitempty"` // Date/time of last cleaning, omit if not applicable
	Notes       string   `json:"notes,omitempty"`        // Any additional notes about the bed
}
