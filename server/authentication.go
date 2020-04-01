package server

import (
	"encoding/json"
	"github.com/JasonSteinberg/timeTicker/responses"
	"log"
	"net/http"

	"github.com/JasonSteinberg/timeTicker/structs"
	"github.com/JasonSteinberg/timeTicker/users"
	"golang.org/x/crypto/bcrypt"
)

// curl -X POST -d '{"email":"count_dooku","password":"iHateWookies"}' http://localhost:8808/signup
func signup(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	var error structs.ServerMessage

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing."
		responses.ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing."
		responses.ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hash)
	err = users.CreateUser(user)

	if err != nil {
		error.Message = err.Error()
		responses.ErrorResponder(w, http.StatusInternalServerError, error)
		return
	}

	HappyMessage := structs.ServerMessage{Message: "Created user successfully"}
	w.Header().Set("Content-Type", "application/json")
	responses.Responder(w, HappyMessage)
}

// curl -X POST -d '{"email":"count_dooku","password":"iHateWookies"}' http://localhost:8808/login
func login(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	var jwToken structs.JWT
	var error structs.ServerMessage

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing."
		responses.ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing."
		responses.ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	password := user.Password
	login, err := users.CheckUser(user, password)

	if err != nil || login == false {
		error.Message = err.Error()
		responses.ErrorResponder(w, http.StatusUnauthorized, error)
		return
	}

	token := GenerateToken(user)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwToken.Token = token

	responses.Responder(w, jwToken)
}
