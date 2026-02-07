package model

type Ward struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	WardType        string      `json:"ward_type"`
	NumberOfBeds    int         `json:"number_of_beds"`
	OccupiedBeds    int         `json:"occupied_beds"`
	AvailableBeds   int         `json:"available_beds"`
	WardStaffing    Staffing    `json:"staffing_levels"`
	WardPolicies    []string    `json:"ward_policies"`
	WardLayout      WardLayout  `json:"ward_layout"`
	WardLocation    Location    `json:"location"`
	WardCleanliness Cleanliness `json:"cleanliness_level"`
}

type WardLayout struct {
	Beds []Bed `json:"beds"`
}

type Staffing struct {
	Doctors int `json:"doctors"`
	Nurses  int `json:"nurses"`
	Others  int `json:"others"` // Other staff members
}

type Cleanliness struct {
	RequiredLevel string `json:"required_level"`
	Level         string `json:"level"` // e.g., very High, High, Medium, Low
	LastCleaned   string `json:"last_cleaned"`
}
