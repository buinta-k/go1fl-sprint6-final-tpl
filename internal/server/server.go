package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger log.Logger
	server *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandleOne)
	mux.HandleFunc("/upload", handlers.HandleTwo)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: *logger,
		server: httpServer,
	}

}

func (s *Server) Start() error {
	s.logger.Println("Сервер успешно запущен")
	return s.server.ListenAndServe()
}
