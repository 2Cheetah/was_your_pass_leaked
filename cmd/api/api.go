package api

import (
	"log/slog"
	"net/http"

	internalhandlers "github.com/2Cheetah/was_your_pass_leaked/internal/handlers"
	sharedhandlers "github.com/2Cheetah/was_your_pass_leaked/pkg/handlers"
)

type APIServer struct {
	Addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		Addr: addr,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/ping", sharedhandlers.PingHandler)
	mux.HandleFunc("POST /api/v1/isLeaked", internalhandlers.IsLeakedHandler)

	slog.Info("Starting server:", "host", s.Addr)
	return http.ListenAndServe(s.Addr, mux)
}
