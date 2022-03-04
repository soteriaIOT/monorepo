package auth

import (
	"github.com/arora-aditya/monorepo/application-server/graph/model"

	"github.com/golang-jwt/jwt"
)

var issuer = []byte("soteria")

// DecodeJwt decode jwt
func DecodeJwt(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return issuer, nil
	})
}

// GenerateJwt create jwt
func GenerateJwt(username string, expiredAt int64) string {
	claims := model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    string(issuer),
		},
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(issuer)

	return signedToken
}