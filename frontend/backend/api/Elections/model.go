package elections

type Election struct {
	ElectionID string `json:"electionID"`
	StartDate  string `json:"startDate"`
	EndtDate   string `json:"endDate"`
}
