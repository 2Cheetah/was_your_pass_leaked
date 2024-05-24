package main

import (
	"log/slog"

	"github.com/2Cheetah/was_your_pass_leaked/cmd/api"
)

func main() {
	server := api.NewAPIServer("localhost:8080")
	if err := server.Run(); err != nil {
		slog.Error("Error, while trying to start server:", "host", server.Addr)
		panic(err)
	}
}
