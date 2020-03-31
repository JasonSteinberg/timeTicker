package server

import (
	"errors"
	"github.com/JasonSteinberg/timeTicker/structs"
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
			return []byte(ultraSecret), nil
		})

		if error != nil {
			errorObject.Message = error.Error()
			ErrorResponder(w, http.StatusUnauthorized, errorObject)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			errorObject.Message = error.Error()
			ErrorResponder(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}
