package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func PingHandler(w http.ResponseWriter, _ *http.Request) {
	respMap := map[string]string{"ping": "pong"}
	respJson, err := json.Marshal(respMap)
	if err != nil {
		slog.Warn(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}
