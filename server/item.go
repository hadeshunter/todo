package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) createItem(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body struct {
		Title string `json:"title"`
	}

	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Request body is invalid")
		return
	}

	if item, err := server.db.CreateItem(body.Title); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to create item")
	} else {
		respondWithJSON(w, http.StatusCreated, item)
	}
}

func (server *Server) deleteItem(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	} else if item, err := server.db.DeleteItem(uint(id)); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	} else {
		respondWithJSON(w, http.StatusOK, &item)
	}
}

func (server *Server) toggleItem(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	} else if item, err := server.db.ToggleItem(uint(id)); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	} else {
		respondWithJSON(w, http.StatusOK, &item)
	}
}

func (server *Server) completeItem(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	} else if item, err := server.db.CompleteItem(uint(id)); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	} else {
		respondWithJSON(w, http.StatusOK, &item)
	}
}

func (server *Server) listAllItems(w http.ResponseWriter, r *http.Request) {
	if items, err := server.db.ListAllItems(); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to retrieve items")
	} else {
		respondWithJSON(w, http.StatusOK, items)
	}
}
