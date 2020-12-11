package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/hadeshunter/todo/database"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

// Server will handle all request via rest API
type Server struct {
	router  *mux.Router
	db      *database.Database
	running bool
}

// New return instance of server
func New() *Server {
	server := &Server{
		db: database.New(os.Getenv("DATABASE_URL")),
	}
	server.initializeRoutes()
	return server
}

func (server *Server) initializeRoutes() {
	server.router = mux.NewRouter()
	server.router.Use(authHandler)
	server.router.HandleFunc("/", server.sayHello).Methods("GET")
	server.router.HandleFunc("/login", server.handleLogin).Methods("GET")
	server.router.HandleFunc("/user/create", server.createUser).Methods("POST")
	server.router.HandleFunc("/user/all", server.getAllUsers).Methods("GET")
	server.router.HandleFunc("/item/create", server.createItem).Methods("POST")
	server.router.HandleFunc("/item/{id:[0-9]+}/delete", server.deleteItem).Methods("DELETE")
	server.router.HandleFunc("/item/{id:[0-9]+}/complete", server.completeItem).Methods("PUT")
	server.router.HandleFunc("/item/{id:[0-9]+}/toggle", server.toggleItem).Methods("PUT")
	server.router.HandleFunc("/item/all", server.listAllItems).Methods("GET")
	server.router.HandleFunc("/unit/all", server.listAllUnits).Methods("GET")
}

func (server *Server) sayHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Wellcome to backend API by Huynh Minh Tri"))
}

// Start the server
func (server *Server) Start(url string) {
	server.running = true

	// TODO: Find the way to monitor better
	// go server.monitorPayment()

	log.Println("Server is ready at", url)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "DELETE", "PUT", "GET", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(server.router)
	http.ListenAndServe(url, handler)
}

// Stop the server
func (server *Server) Stop() {
	log.Println("Try to shutdown gracefully...")
	server.running = false
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}

// GetRouter ..
func (server *Server) GetRouter() *mux.Router {
	return server.router
}
