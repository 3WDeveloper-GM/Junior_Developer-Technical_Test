package app

type Jwt interface {
	VerifyToken(tokenString string) error
	CreateToken(username string) (string, error)
}

