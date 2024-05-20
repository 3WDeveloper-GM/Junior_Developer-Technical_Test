package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
)

func GenerateNewToken(usedID int, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserId: usedID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	tokenHash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = tokenHash[:]

	return token, nil
}

func ValidateTokenLength(v validate, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "tokenPlaintext", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "tokenPlaintext", "token must be exactly 26 characters long")
}
