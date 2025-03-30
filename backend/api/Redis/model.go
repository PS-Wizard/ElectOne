package redis

type vote struct {
	CandidateID string `json:"candidateID"`
	ElectionID  string `json:"electionID"`
	// TODO: Perhaps add additional fields such as password for confirmation
}
