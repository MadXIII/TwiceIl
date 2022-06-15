package handler

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Run(address string, handlers http.Handler) error {
	s.server = &http.Server{
		Addr:           address,
		Handler:        handlers,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.server.ListenAndServe()
}
