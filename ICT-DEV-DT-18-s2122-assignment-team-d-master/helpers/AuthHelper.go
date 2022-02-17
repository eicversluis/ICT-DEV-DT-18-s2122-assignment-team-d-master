package helpers

import (
	"context"
	"net/http"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" {
			username, err := r.Cookie("username")
			if err != nil || username.Value == "" {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			} else {
				ctx := context.WithValue(r.Context(), "username", username.Value)
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
