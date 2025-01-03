package entities

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

type UserNotPassword struct {
	ID    int
	Name  string
	Email string
}
