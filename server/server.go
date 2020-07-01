package server

import (
	"encoding/json"
	"net/http"
	"log"
	"os"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/hadeshunter/todo/database"
)
// Server api
type Server struct{
	db *database.Database
}
// New server
func New() *Server{
	server := &Server{
		db: database.New(os.Getenv("DATABASE_URL")),
	}
	return server
}
// Start server
func (server *Server) Start(url string) {
	router := mux.NewRouter()
	router.HandleFunc("/", server.sayHello).Methods("GET")
	log.Println("Server is starting at", url)

	//------ Handle here------//
	router.HandleFunc("/item/create", server.createItem).Methods("POST")
	router.HandleFunc("/item/complete", server.completeItem).Methods("PUT")
	router.HandleFunc("/item/all", server.listAllItems).Methods("GET")
	//------------------------//

	corsPolicy := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{"POST", "DELETE", "PUT", "GET", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
	handlerWithCors := corsPolicy.Handler(router)
	http.ListenAndServe(url, handlerWithCors)
}

func (server *Server) sayHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Trí"))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error":message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}