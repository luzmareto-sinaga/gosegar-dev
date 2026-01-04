package response

// mengacu pada File SingIn pada : https://docs.google.com/spreadsheets/d/1cvyRqBO8yG0wWNkU3lO6SrqCPWWm-os_81dFAliKuPA/edit?gid=1763570654#gid=1763570654
type SignInResponse struct {
	AccessToken string `json:"access_token"`
	Role        string `json:"role"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Lat         string `json:"latitude"`
	Lng         string `json:"longitude"`
}
