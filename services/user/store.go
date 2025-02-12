package user

import (
	"database/sql"
	"fmt"

	"github.com/ImArnav19/ecom/models"
)

type Store struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*models.User, error) {

	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	u := new(models.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*models.User, error) {

	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil

}

func (s *Store) CreateUser(user models.User) error {
	_, err := s.db.Exec("INSERT INTO users(firstName,lastName,email,password) VALUES(?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
