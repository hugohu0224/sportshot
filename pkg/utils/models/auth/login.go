package auth

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
