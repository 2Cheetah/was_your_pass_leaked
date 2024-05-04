package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	fmt.Println("Starting server")

	startServer()
}

func startServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", pingHandler)
	mux.HandleFunc("POST /isLeaked", isLeakedHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err.Error())
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	respMap := map[string]string{"ping": "pong"}
	respJson, err := json.Marshal(respMap)
	if err != nil {
		slog.Warn(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func isLeakedHandler(w http.ResponseWriter, r *http.Request) {
	respMap := map[string]bool{"isLeaked": true}
	respJson, err := json.Marshal(respMap)
	if err != nil {
		slog.Warn(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}
