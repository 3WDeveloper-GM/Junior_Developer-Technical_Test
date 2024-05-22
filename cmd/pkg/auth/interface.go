package auth

import "time"

type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserId    int64       `json:"-"`
	Expiry    time.Time `json:"expiracion"`
	Scope     string    `json:"-"`
}

type validate interface {
	Valid() bool
	Check(bool, string, string)
	AddErrorKey(string, string)
}
