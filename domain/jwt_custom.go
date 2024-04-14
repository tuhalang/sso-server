package domain

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	SessionId string `json:"sessionId"`
	jwt.RegisteredClaims
}
