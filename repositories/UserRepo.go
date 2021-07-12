package repositories

import (
	"backend/models"
	"database/sql"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo ..
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// FindByID ..
func (r *UserRepo) FindByID(ID int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("select userid, username from users where userid = $1", ID).Scan(&user.Userid, &user.Username)
	return &user, err
}

func (r *UserRepo) FindByUsername(Username string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("select userid, username, hash from users where username = $1", Username).Scan(&user.Userid, &user.Username, &user.Hash)
	return &user, err
}

// Save ..
func (r *UserRepo) Save(user *models.User) error {
	err := r.db.QueryRow("insert into users(username, hash) values($1, $2)", user.Username, user.Hash).Err()
	return err
}
