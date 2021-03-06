package middleware

import (
	"context"
	"errors"
	"github.com/JasonSteinberg/timeTicker/responses"
	"github.com/JasonSteinberg/timeTicker/structs"
	"github.com/JasonSteinberg/timeTicker/users"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

func ProtectedMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject structs.ServerMessage
		authToken := r.Header.Get("Authorization")

		token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("There was an error!")
			}
			return []byte(structs.UltraSecret), nil
		})

		if error != nil {
			errorObject.Message = error.Error()
			responses.ErrorResponder(w, http.StatusUnauthorized, errorObject)
			return
		}

		if token.Valid {
			email := token.Claims.(jwt.MapClaims)["email"].(string)
			user := users.FillUser(email)
			ctx := context.WithValue(r.Context(), "User", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			errorObject.Message = error.Error()
			responses.ErrorResponder(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}
