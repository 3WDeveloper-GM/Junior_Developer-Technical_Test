package app

type Jwt interface {
	VerifyToken(tokenString string) (string, error)
	CreateToken(username string, email string) (string, error)
}

