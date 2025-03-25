package citizens

type citizen struct {
    CitizenID string `json:"citizenID"`
	Name      string `json:"fullName"`
	DOB       string `json:"dateOfBirth"`
	POR       string `json:"placeOfResidence"`
}

