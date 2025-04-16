package candidates

type Candidate struct {
	CandidateID string `json:"candidateID"`
	CitizenID   string `json:"citizenID"`
	Post        string `json:"post"`
	ElectionID  string `json:"electionID"`
	Group       string `json:"group"`
}
