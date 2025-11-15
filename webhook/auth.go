package webhook

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// BasicAuthValidator creates a RequestValidator function that validates Basic Authentication
// It accepts a function that takes username and password and returns whether they are valid
func BasicAuthValidator(validateCredentials func(username, password string) bool) func(*http.Request) error {
	return func(r *http.Request) error {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			return errors.New("missing Authorization header")
		}

		// Check if it's Basic auth
		if !strings.HasPrefix(auth, "Basic ") {
			return errors.New("invalid authorization scheme, expected Basic")
		}

		// Decode the credentials
		encoded := strings.TrimPrefix(auth, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			return fmt.Errorf("invalid base64 encoding: %w", err)
		}

		// Split username:password
		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			return errors.New("invalid credentials format")
		}

		username := credentials[0]
		password := credentials[1]

		// Validate credentials using the provided function
		if !validateCredentials(username, password) {
			return errors.New("invalid credentials")
		}

		return nil
	}
}

// BasicAuthErrorHandler creates an OnError handler that returns 401 Unauthorized for auth errors
func BasicAuthErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}
	errMsg := err.Error()
	if strings.Contains(errMsg, "Authorization") ||
		strings.Contains(errMsg, "credentials") ||
		strings.Contains(errMsg, "invalid") {
		w.Header().Set("WWW-Authenticate", `Basic realm="Webhook"`)
		http.Error(w, errMsg, http.StatusUnauthorized)
		return
	}
	http.Error(w, errMsg, http.StatusInternalServerError)
}
