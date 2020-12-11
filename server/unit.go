package server

import "net/http"

func (server *Server) listAllUnits(w http.ResponseWriter, r *http.Request) {
	if units, err := server.db.ListAllUnits(); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to retrieve units")
	} else {
		respondWithJSON(w, http.StatusOK, units)
	}
}
