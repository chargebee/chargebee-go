package webhook

import (
	"errors"
	"net/http"
)

func BasicAuthValidator(validateCredentials func(username, password string) bool) func(*http.Request) error {
	return func(r *http.Request) error {
		username, password, ok := r.BasicAuth()
		if !ok {
			return errors.New("invalid basic authorization header")
		}

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
		return
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="Webhook"`)
	http.Error(w, "unauthorized", http.StatusUnauthorized)
}
