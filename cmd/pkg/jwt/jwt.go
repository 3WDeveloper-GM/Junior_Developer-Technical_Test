package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type customClaims struct {
	User  string `json:"user"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

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

func (jw *jwtToken) CreateToken(username string, email string) (string, error) {
	claims := customClaims{
		username,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "billings",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jw.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jw *jwtToken) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jw.secretKey, nil
		})

	claims, ok := token.Claims.(*customClaims)

	if ok {
		log.Info().Msg(claims.Email)
	} else {
		return "", errInvalidToken
	}

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errInvalidToken
	}

	return claims.Email, nil
}
