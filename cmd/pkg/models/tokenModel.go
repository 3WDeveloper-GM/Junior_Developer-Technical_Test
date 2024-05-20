package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/auth"
)

type tokenModel struct {
	DB *sql.DB
}

func (tk *tokenModel) Insert(token *auth.Token) error {
	query := `
    INSERT INTO tokens (hash, user_id, expiry, scope)
    VALUES ($1, $2, $3, $4)
  `
	args := []interface{}{token.Hash, token.UserId, token.Expiry, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := tk.DB.ExecContext(ctx, query, args...)
	return err
}

func (tk *tokenModel) New(user int, ttl time.Duration, scope string) (*auth.Token, error) {
	token, err := auth.GenerateNewToken(user, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = tk.Insert(token)
	return token, err
}

func (tk *tokenModel) DeleteAllTokensFromUser(scope string, userID int) error {
	query := `
      DELETE FROM tokens
      WHERE scope = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := tk.DB.ExecContext(ctx, query, scope, userID)
	return err
}
