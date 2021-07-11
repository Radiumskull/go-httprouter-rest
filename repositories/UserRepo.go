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
	err := r.db.QueryRow("select username from users where userid = $1", ID).Scan(&user.Username)
	return &user, err
}

// Save ..
func (r *UserRepo) Save(user *models.User) error {
	return nil
}
