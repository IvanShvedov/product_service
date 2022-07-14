package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"product-service/internal/config"
	"product-service/internal/services"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	config  *config.Config
	router  *mux.Router
	service services.Service
}

func New(cfg *config.Config, db *sql.DB) *Server {
	return &Server{
		config:  cfg,
		router:  mux.NewRouter(),
		service: services.NewBaseService(db),
	}
}

func (s *Server) Start() error {
	s.configureRouter()
	return http.ListenAndServe(s.config.Listen.BindIP+s.config.Listen.Port, s.router)
}

func (s *Server) configureRouter() {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/products", s.create_product()).Methods("POST")
	s.router.HandleFunc("/products", s.update_product()).Methods("PUT")
	s.router.HandleFunc("/products", s.delete_product()).Methods("DELETE")
	s.router.HandleFunc("/products", s.fetch_all_product()).Methods("GET")
	s.router.HandleFunc("/products/{id}", s.fetch_one_product()).Methods("GET")
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
