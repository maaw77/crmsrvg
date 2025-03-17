package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/maaw77/crmsrvg/internal/auth"
	"github.com/maaw77/crmsrvg/internal/handlers/v1"
)

func subHaldler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("subHaldler Serving:", r.URL.Path, " from", r.Host)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "it's OK")
}
func TestAuthMiddlewareBadToken(t *testing.T) {
	token, _ := auth.GetToken("Konrad", time.Now().Add(2*time.Second).Unix())
	handler := AuthMiddleware(http.HandlerFunc(subHaldler))

	r := httptest.NewRequest("GET", "/gsm", nil)
	r.Header.Set("Authokization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	// t.Log(w.Code)
	// Check the HTTP status code is what we expect.
	if w.Code != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusOK)
		return
	}

	var errA handlers.ErrorMessage
	json.NewDecoder(w.Body).Decode(&errA)
	// t.Log(errA.Details)
	if errA.Details != "token is not provided" {
		t.Errorf(`%s != "token is not provided"`, errA)
	}

	r = httptest.NewRequest("GET", "/gsm", nil)
	// r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w = httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	// t.Log(w.Code)
	// Check the HTTP status code is what we expect.
	if w.Code != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusOK)
		return
	}

	json.NewDecoder(w.Body).Decode(&errA)
	// t.Log(errA.Details)
	if errA.Details != "token is not provided" {
		t.Errorf(`%s != "token is not provided"`, errA)
	}

	r = httptest.NewRequest("GET", "/gsm", nil)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "token"))
	w = httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	// t.Log(w.Code)
	// Check the HTTP status code is what we expect.
	if w.Code != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusOK)
		return
	}

	json.NewDecoder(w.Body).Decode(&errA)
	// t.Log(errA.Details)
	if errA.Details != "that's not even a token" {
		t.Errorf(`%s != "that's not even a token"`, errA)
	}

	time.Sleep(3 * time.Second)

	r = httptest.NewRequest("GET", "/gsm", nil)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	w = httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	// t.Log(w.Code)
	// Check the HTTP status code is what we expect.
	if w.Code != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusOK)
		return
	}

	json.NewDecoder(w.Body).Decode(&errA)
	// t.Log(errA.Details)
	if errA.Details != "token is expired" {
		t.Errorf(`%s != "token is expired"`, errA)
	}

}

func TestAuthMiddleware(t *testing.T) {
	token, _ := auth.GetToken("Konrad", time.Now().Add(2*time.Second).Unix())
	handler := AuthMiddleware(http.HandlerFunc(subHaldler))

	r := httptest.NewRequest("GET", "/gsm", nil)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	// t.Log(w.Code)
	// Check the HTTP status code is what we expect.
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusOK)
		return
	}

	bodySring := w.Body.String()
	if bodySring != "it's OK" {
		t.Errorf(`%s != "it's OK"`, bodySring)
	}
}
