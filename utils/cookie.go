package utils

import "net/http"

// ParseCookie finds cookie in request if present
func ParseCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}
