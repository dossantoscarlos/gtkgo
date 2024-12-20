package dto

type AuthRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthDTO struct {
	Username string `json:"nome"`
	Email    string `json:"email"`
}

type AuthResponseDTO struct {
	Token string `json:"token"`
}
