package server

import (
	"encoding/json"
	"net/http"
)

func (server *Server) createItem(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var body struct {
		Title string `json:"title"`
	}

	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Request body is invalid")
		return
	}

	if item, err := server.db.CreateItem(body.Title); err!= nil{
		respondWithError(w, http.StatusInternalServerError, "Unable to create item")
	} else {
		respondWithJSON(w, http.StatusCreated, item)
	}
}

func (server *Server) completeItem(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body struct {
		ID uint `json:"id"`
	}

	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Request body is invalid")
		return
	}

	if err := server.db.CompleteItem(body.ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to complete item")
	} else {
		respondWithJSON(w, http.StatusOK, "")
	}
}

func (server *Server) listAllItems(w http.ResponseWriter, r *http.Request) {
	if items, err := server.db.ListAllItems(); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to retrieve items")
	} else {
		respondWithJSON(w, http.StatusOK, items)
	}
}
