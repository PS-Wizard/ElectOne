package citizens

type Citizen struct {
    CitizenID string `json:"citizenID"`
	Name      string `json:"fullName"`
	DOB       string `json:"dateOfBirth"`
	POR       string `json:"placeOfResidence"`
}

