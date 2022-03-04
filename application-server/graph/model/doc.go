package model

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}