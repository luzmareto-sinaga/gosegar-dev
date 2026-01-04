package response

// mengacu pada File SingIn pada : https://docs.google.com/spreadsheets/d/1cvyRqBO8yG0wWNkU3lO6SrqCPWWm-os_81dFAliKuPA/edit?gid=1763570654#gid=1763570654
type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
