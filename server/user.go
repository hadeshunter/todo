package server

import (
	"net/http"
	"encoding/json"
)

func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body struct {
		Name		string		`json:"name"`
		Phone		string		`json:"phone"`
		Email		string		`json:"email"`
	}

	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Request body is invalid")
		return
	}

	if user, err := server.db.CreateUser(body.Name, body.Phone, body.Email); err!= nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to create item")
	} else {
		respondWithJSON(w, http.StatusCreated, user)
	}
}

func (server *Server) getAllUsers(w http.ResponseWriter, r *http.Request) {
	if users, err := server.db.GetAllUsers(); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to retrieve users")
	} else {
		respondWithJSON(w, http.StatusOK, users)
	}
}