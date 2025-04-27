package appeals

type Appeal struct {
	AppealID int `json:"appealID"`

	//Citizen details
	CitizenID string `json:"citizenID"`
	Name      string `json:"fullName"`
	DOB       string `json:"dateOfBirth"`
	POR       string `json:"placeOfResidence"`

	// User details
	Password string   `json:"password"`
	Phone    string   `json:"phoneNumber"`
	Tag      string   `json:"tags"`
	Photos   []string `json:"photos"` // Assuming photos are stored as a JSON array or comma-separated values
}
