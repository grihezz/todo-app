package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:                         ":" + port,
		Handler:                      handler,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  10 * time.Second,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  0,
		MaxHeaderBytes:               1 << 20,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
