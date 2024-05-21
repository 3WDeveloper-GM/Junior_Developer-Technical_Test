package models

import (
	"context"
	"crypto/sha256"
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
    INSERT INTO users (provider_id, name, email, password_hash, activated)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, created_at, version`

	args := []interface{}{user.ProviderID, user.Name, user.Email, user.Password.Hash, user.Activated}
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

func (m *userModel) GetForToken(tokenScope string, tokenPlaintext string) (*domain.Users, error) {
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	query := `
    SELECT users.id, users.created_at, users.name, users.email, users.password_hash, users.activated, users.version
    FROM users
    INNER JOIN tokens
    ON users.id = tokens.user_id
    WHERE tokens.hash = $1
    AND tokens.scope = $2
    AND tokens.expiry > $3`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{tokenHash[:], tokenScope, time.Now().Format(time.DateOnly)}

	var user domain.Users
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.SysID,
		&user.Created_at,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Activated,
		&user.Version,
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
