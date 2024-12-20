package dto

type UserDTO struct {
	Name     string `json:"nome"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDtoResponse struct {
	Name  string `json:"nome"`
	Email string `json:"email"`
}
