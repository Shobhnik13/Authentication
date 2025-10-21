package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthProtect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		// const decoded = jwt.verify(tokenString, env secret ) in node.js

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// here we need to extract claims , unlike node.js we will first get token here then we need to get the claims using that token then we will attach but in node.js u verify and u dont need to do further like extract claims and all the verify return u extracted data
		// here we first parse the token then from the parsed token we extract claims
		// attaching user info to request context
		// req.user = decoded in node.js
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized: Invalid token claims", http.StatusUnauthorized)
			return
		}

		// its just like req.user = decoded in node.js
		// we are getting a context of request and attaching the user info to that context
		ctx := context.WithValue(r.Context(), "user", claims)

		// replace the old context with the new one
		// basically overriding the old context with the new one which has user info
		// next() in node.js
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
