package utils

import (
	"fmt"
	"os"
	"time"
	"vk-test/pkg/database/mongodb/repository"

	"github.com/golang-jwt/jwt/v5"
)

const adminEmail string = "chingizkhan@gmail.com"

// Function that handles Token generation using a secret key
func GenerateAllTokens(email string, uid string) (signedToken string, signedRefreshToken string, err error) {

	/*err = godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}*/
	var Admin bool
	if email == adminEmail {
		Admin = true
	}

	var SECRET_KEY string = os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":   email,
		"isAdmin": Admin,
		"Uid":     uid,
		"exp":     time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
	})

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":   email,
		"isAdmin": Admin,
		"Uid":     uid,
		"exp":     time.Now().Local().Add(time.Hour * time.Duration(169)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return
	}
	refreshToken, err := refresh.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return
	}

	return tokenString, refreshToken, err
}

// Function that validates token and returns its claims
func ValidateToken(tokenString string) (jwt.MapClaims, error) {

	var SECRET_KEY string = os.Getenv("SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

// Function that updates the token for a given user
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) (string, string, error) {

	_, err := ValidateToken(signedRefreshToken)
	if err != nil {
		return "", "", err
	}

	err = repository.UpdateTokens(signedToken, signedRefreshToken, userId)

	if err != nil {
		return "", "", err
	}

	return signedToken, signedRefreshToken, nil

}
