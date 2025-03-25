package users

type user struct {
	UserID    int    `json:"userID"`
	CitizenID string `json:"citizenID"`
	Password  string `json:"password"`
}
