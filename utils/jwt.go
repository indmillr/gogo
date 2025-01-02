package utils

import (
	"errors"
	// "fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretSquirrel"

func GenerateToken(email string, userId int64) (string, error) {
	// ----- userId was consistently logging as 0 because in user.go, the definition of ValidateCredentials() was not using a pointer, so it was using a Copy of the User. Changing it to *User for ValidateCredentials() seems to have solved this.
	// fmt.Println(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// ----- This is a type-checking syntax...
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected Signing Method.")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token.")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid Token!")
	}

	// ----- More type-checking...
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims.")
	}

	// // ----- More type-checking...
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}