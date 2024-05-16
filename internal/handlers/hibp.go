package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/2Cheetah/was_your_pass_leaked/pkg/helper"
)

func IsLeakedHandler(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	isLeaked := helper.CheckLeaked(password)

	respMap := map[string]bool{"isLeaked": isLeaked}
	respJson, err := json.Marshal(respMap)
	if err != nil {
		slog.Warn(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}
