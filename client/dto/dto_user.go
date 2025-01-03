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

type UserCreateResponse struct {
	ID int `json:"id"`
}

type UserIdNameResponse struct {
	ID   int    `json:"id"`
	Name string `json:"nome"`
}
