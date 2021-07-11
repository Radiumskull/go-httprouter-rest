package models

// User ..
type User struct {
	Userid   int
	Username string
	Hash     string
}

// UserRepository ..
type UserRepository interface {
	FindByID(ID int) (*User, error)
	Save(user *User) error
}
