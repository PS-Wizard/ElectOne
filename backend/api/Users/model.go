package users

type User struct {
	UserID    int    `json:"userID"`
	CitizenID string `json:"citizenID"`
	Password  string `json:"password"`
	Phone     string `json:"phoneNumber"`
	Tag       string `json:"tags"`
	OTP       string `json:"otp"`
	Photos    string `json:"photos"`
}
