package models

import (
	"database/sql"
	"errors"
)

type AppModels struct {
	Bills   *billModel
	Users   *userModel
	Tokens  *tokenModel
	Permits *permitModel
}

var ErrNoRows = errors.New("record Not found, either bill or user does not exist")

func InitializeAppModels(db *sql.DB) *AppModels {
	bill := &billModel{
		DB: db,
	}
	user := &userModel{
		DB: db,
	}
	token := &tokenModel{
		DB: db,
	}
	permits := &permitModel{
		DB: db,
	}
	return &AppModels{
		Bills:   bill,
		Users:   user,
		Tokens:  token,
		Permits: permits,
	}
}
