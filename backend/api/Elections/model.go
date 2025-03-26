package elections

type Election struct {
	ElectionID         string `json:"electionID"`
	Title              string `json:"title"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"endDate"`
	VotingRestrictions string `json:"votingRestriction"`
}
