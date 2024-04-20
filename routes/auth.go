package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"sailing-seas/core"
	"strings"
	"time"

	"github.com/google/uuid"
)

var cookies = map[string]time.Time{}

func IsValidToken(token string) bool {
	return time.Now().Before(cookies[token])
}

func IsAuthenticated(r *http.Request) bool {
	cookie, _ := r.Cookie("token")
	return cookie != nil && IsValidToken(cookie.Value)
}

func NeedsAuthentication(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthenticated(r) {
			next := url.URL{
				Path:     r.URL.Path,
				RawQuery: r.URL.RawQuery,
			}
			url := fmt.Sprintf("/login?next=%s", url.QueryEscape(next.String()))
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			return
		}
		handler(w, r)
	}
}

func Authenticate(app *core.App, w http.ResponseWriter, username string, password string) bool {
	if username != app.Env.Username || password != app.Env.Password {
		return false
	}
	token := uuid.NewString()
	expiresAt := time.Now().Add(5 * time.Minute)
	cookies[token] = expiresAt
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiresAt,
	})
	return true
}

func redirect(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	if next == "" || !strings.HasPrefix(next, "/") || next == "/login" {
		next = "/"
	}
	http.Redirect(w, r, next, http.StatusTemporaryRedirect)
}
