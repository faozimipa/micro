package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type token struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

type TokenResponse struct {
	AccessToken token `json:"access_token"`
}

func GenerateToken(userID uuid.UUID) token {
	claims := token{
			UserID: userID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				Audience:  "http://krakend:5000",
				Issuer:    "http://identity-service:8081",
			},
	}
	tokenize := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := []byte("my_secret_key")
	tokenString, _ := tokenize.SignedString(mySigningKey)
	fmt.Println(tokenString)


	return token{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Audience:  "http://krakend:5000",
			Issuer:    "http://identity-service:8081",
		},
	}

	// claims := token{
	// 	UserID: userID,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	// 		Audience:  "http://krakend:5000",
	// 		Issuer:    "http://identity-service:8081",
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// mySigningKey := []byte("my_secret_key")
	// tokenString, _ := token.SignedString(mySigningKey)

	// return tokenString

	// KrakenD signs the access token. Uncomment to sign in identity service.
	// token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	// mySigningKey := []byte("my_secret_key")
	// tokenString, err := token.SignedString(mySigningKey)
	// fmt.Println(tokenString)
}
