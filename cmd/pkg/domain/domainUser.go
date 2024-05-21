package domain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/handlers/validator"
	"golang.org/x/crypto/bcrypt"
)

var AnonUser = &Users{}

func NewPassword() *passWord {
	return &passWord{}
}

func (p *passWord) Set(plaintext string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return err
	}

	p.Plaintext = &plaintext
	p.Hash = hash

	return nil
}

func (p *passWord) Matches(plaintext string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintext))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			fmt.Println("got here")
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateEmail(v validate, email string) {
	v.Check(email == strings.ToLower(email), "correoUsuario", "must be lowecase")

	v.Check(email != "", "correoUsuario", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "correoUsuario", "must be a valid email address")
}

func ValidatePasswordFromPlaintext(v validate, plaintext string) {
	v.Check(plaintext != "", "contraUsuario", "must be provided")
	v.Check(len(plaintext) >= 8, "contraUsuario", "must be at least 8 characters long")
	v.Check(len(plaintext) <= 50, "contraUsuario", "must be less than 50 characters long")
}

func (usr *Users) ValidateUser(v validate) bool {
	v.Check(usr.ProviderID != "", "idProveedor", "must be provided")
	v.Check(len(usr.ProviderID) == 36, "idProveedor", "must be a valid uuid signature")

	v.Check(usr.Name != "", "nombreUsuario", "must be provided")
	v.Check(len(usr.Name) <= 500, "nombreUsuario", "must not be more than 500 bytes long")

	ValidateEmail(v, usr.Email)

	if usr.Password.Plaintext != nil {
		ValidatePasswordFromPlaintext(v, *usr.Password.Plaintext)
	}

	if usr.Password.Hash == nil {
		panic("missing password hash for usr")
	}

	return v.Valid()
}

func (usr *Users) IsAnonymous() bool {
	return usr == AnonUser
}
