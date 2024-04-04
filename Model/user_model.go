package model

type User struct {
	ID    int
	Email string
	Role  string
}

func NewUser(id int, email, role string) *User {
	return &User{
		ID:    id,
		Email: email,
		Role:  role,
	}
}
