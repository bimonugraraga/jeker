package utils

import (
	r "github.com/bdn/jeker/dto/responses"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Params r.JWTTokenFormat
	jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key")

func SignJWT(params r.JWTTokenFormat) (string, error) {
	claims := &Claims{
		Params: params,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(accessToken string) (r.JWTTokenFormat, error) {
	claims := &Claims{}
	var payload r.JWTTokenFormat
	tkn, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil {
		return payload, err
	}
	if !tkn.Valid {
		return payload, err
	}
	payload = r.JWTTokenFormat{
		Email:    claims.Params.Email,
		Username: claims.Params.Username,
	}

	return payload, nil
}
