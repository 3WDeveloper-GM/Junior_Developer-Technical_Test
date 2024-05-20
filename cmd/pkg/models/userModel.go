package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/domain"
)

type userModel struct {
	DB *sql.DB
}

var ErrDuplicateMail = errors.New("duplicate email")

func (m *userModel) Create(user *domain.Users) error {
	query := `
    INSERT INTO users (name, email, password_hash, activated)
    VALUES ($1, $2, $3, $4)
    RETURNING id, created_at, version`

	args := []interface{}{user.Name, user.Email, user.Password.Hash, user.Activated}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.SysID, &user.Created_at, &user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateMail
		default:
			return err
		}
	}

	return nil
}

func (m *userModel) Fetch(email string) (*domain.Users, error) {
	query := `
    SELECT id, created_at, name, email, password_hash, activated
    FROM users
    WHERE email = $1
  `

	args := []interface{}{email}

	var user domain.Users

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
  defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.SysID,
		&user.Created_at,
		&user.Name,
		&user.Email,
    &user.Password.Hash,
		&user.Activated,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoRows
		default:
			return nil, err
		}
	}

	return &user, nil
}
