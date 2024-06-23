package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	"github.com/pynezz/pynezzentials/ansi"
	"github.com/pynezz/pynezzentials/fsutil"
)

const (
	exp = 9            // Expires in 9 hours
	iss = "pynezz.dev" // Issuer
)

func getSecretKey() string {
	if exists := fsutil.FileExists(".env"); !exists {
		ansi.PrintError("No .env file found")
		os.Exit(1)
	}

	// Read the secret key from the .env file
	envFile, err := fsutil.GetFileContent(".env")
	if err != nil {
		ansi.PrintError(err.Error())
		os.Exit(1)
	}

	secretKey := strings.Split(envFile, "=")[1]
	fmt.Println("Secret key: ", secretKey)

	secretKey = strings.Trim(secretKey, "\"") // Remove any quotes from the secret key
	fmt.Println("Secret key trimmed: ", secretKey)

	return secretKey
}

func VerifyJWTToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		signing := token.Header["alg"]
		fmt.Printf("%v\n", token)
		fmt.Println("Signing method: ", signing)

		exp := token.Claims.(jwt.MapClaims)["exp"].(float64)
		debugExp := fmt.Sprintf("%f", exp)
		ansi.PrintDebug(fmt.Sprintf("Token expires at: %s", debugExp))

		timeNow := time.Now().Unix()
		fmt.Println("Time now: ", timeNow)

		if exp < float64(timeNow) {
			fmt.Println("Token has expired")
			return nil, echo.ErrUnauthorized
		}

		// Ensure token's signing method matches
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.ErrUnauthorized
		}

		// Return the secret key to the jwt.Parse function
		return []byte(getSecretKey()), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !token.Valid {
		fmt.Println("Token is not valid")
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error getting claims")
		return nil, err
	}

	sub := fmt.Sprintf("%f", claims["sub"].(float64))
	ansi.PrintDebug("Subject(user): " + sub)
	aud := fmt.Sprintf("%s", claims["aud"])
	ansi.PrintDebug("Audience(role): " + aud)

	return token, err
}

func GenerateJWTToken(user models.User, loginTime time.Time) string {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// NB! This might not work. It's a time.Time object
		"exp": loginTime.Add(time.Duration(exp) * time.Hour).Unix(),

		"iss": iss,
		"aud": user.Role,
		"sub": user.UserID,
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tokenString)
	return tokenString
}
