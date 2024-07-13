package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ImArnav19/ecom/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	log.Println("Server is running on", s.addr)

	store := user.NewUserStore(s.db)

	userHandler := user.NewUserHandler(store)
	userHandler.MakeRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
