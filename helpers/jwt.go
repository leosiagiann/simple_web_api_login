package helpers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "simplewebapilogin"

func GenerateJwtToken(userId uint, email string) string {
	claims := jwt.MapClaims{
		"userId":   userId,
		"username": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := parseToken.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err)
	}

	return token
}
