package utils

// import "github.com/golang-jwt/jwt"
import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	fmt.Println(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretkey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretkey), nil

	})

	if err != nil {
		return 0, errors.New("could not parse Token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claim")
	}

	// email := claims["email"].(string)
	userId, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("invalid type for user ID, expected float64")
	}
	return int64(userId), nil
}
