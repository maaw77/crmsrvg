package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tokenString string

var expectedUsername = "Konrad"

func TestGetToken(t *testing.T) {
	var err error
	timeExp := time.Now().Add(2 * time.Second).Unix()
	tokenString, err = GetToken(expectedUsername, timeExp)
	if err != nil {
		t.Fatalf("%s != nil", err)
	}
	if tokenString == "" {
		t.Fatalf(`tokenString == ""`)
	}
}

func TestVerifyToken(t *testing.T) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		t.Fatalf("%s != nil", err)
	}

	if !token.Valid {
		t.Fatal("token.Valid != true")
	}

	username, err := token.Claims.GetSubject()
	t.Logf("username=%s", username)
	if err != nil {
		t.Fatalf("%s != nil", err)
	}

	if expectedUsername != username {
		t.Errorf("%s != %s", expectedUsername, username)

	}

	expectedIss := "crm-app"
	iss, err := token.Claims.GetIssuer()
	if err != nil {
		t.Fatalf("%s != nil", err)
	}

	if expectedIss != iss {
		t.Errorf("%s != %s", expectedIss, iss)
	}

	time.Sleep(3 * time.Second)

	_, err = VerifyToken(tokenString)

	if !errors.Is(err, jwt.ErrTokenExpired) {
		t.Fatalf("%s != %s", err, jwt.ErrTokenExpired)
	}

}
