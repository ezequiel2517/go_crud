package authentication

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Email string
	jwt.StandardClaims
}

var JWTKey = []byte(os.Getenv("MY_VARIABLE"))

func GenerarToken(email string) (string, error) {
	claims := &JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}

func ParsearToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)

	if ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
