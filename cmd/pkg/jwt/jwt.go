package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type jwtToken struct {
	secretKey []byte
}

func NewJwtToken() *jwtToken {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	if secretKey == "" {
		log.Fatal().Msg("secret key not set. Set JWT_SECRET_KEY as an environment variable")
	}

	return &jwtToken{secretKey: []byte(secretKey)}
}

func (jw *jwtToken) CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		})
	tokenString, err := token.SignedString(jw.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jw *jwtToken) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jw.secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errInvalidToken
	}

	return nil
}
