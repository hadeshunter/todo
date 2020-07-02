package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// RequestContextKey is key of request context
type RequestContextKey string

var authHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var phone string
		allowPaths := []string{
			"/",
			"/login",			
			"/user/create",
			"/user/all",
			"/item/create",
			"/item/all",
		}
		for _, path := range allowPaths {
			if r.URL.Path == path {
				next.ServeHTTP(w, r)
				return
			}
		}
		if phone, err = auth(r); err != nil {
			respondWithError(w, http.StatusForbidden, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), RequestContextKey("phone"), phone)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// Token of user
type Token struct {
	jwt.StandardClaims
	Phone string
}

var auth = func(r *http.Request) (string, error) {
	tokenHeader := r.Header.Get("Authorization")
	var token Token
	parseResult, err := jwt.ParseWithClaims(
		tokenHeader,
		&token,
		func(*jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SERVER_SECRET")), nil
		})

	if err == nil && parseResult.Valid {
		return token.Phone, nil
	}

	fmt.Println(err, parseResult)

	return "", errors.New("Not authorized")
}
