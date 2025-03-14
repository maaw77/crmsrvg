package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("1a2c4f")

func GetToken(username string, timeExp int64) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,          // Subject (user identifier)
		"iss": "crm-app",         // Issuer
		"exp": timeExp,           // time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(), // Issued at
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString(mySigningKey)
	return
}

func VerifyToken(tokenString string) (token *jwt.Token, err error) {

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	// switch {
	// case token.Valid:
	// 	return
	// case errors.Is(err, jwt.ErrTokenMalformed):
	// 	fmt.Println("That's not even a token")
	// case errors.Is(err, jwt.ErrTokenSignatureInvalid):
	// 	// Invalid signature
	// 	fmt.Println("Invalid signature")
	// case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
	// 	// Token is either expired or not active yet
	// 	fmt.Println("Timing is everything")
	// default:
	// 	fmt.Println("Couldn't handle this token:", err)
	// }

	return
}
