package models

// User ..
type UserData struct {
	Userid   int
	Username string
	Hash     string
}

// UserRepository ..
type UserDataRepository interface {
	FindByID(ID int) (*UserData, error)
	Save(user *UserData) error
}
