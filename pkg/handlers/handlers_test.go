package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	PingHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code is not 200: %#v", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Wrong response content type: %#v", resp.Header.Get("Content-Type"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error while reading response body: %#v", err)
	}

	var expected = map[string]string{"ping": "pong"}
	var bodyData map[string]string
	if err := json.Unmarshal(body, &bodyData); err != nil {
		t.Errorf("Couldn't unmarshall response body: %#v", err)
	}
	if bodyData["ping"] != expected["ping"] {
		t.Errorf("Reponse body data is wrong: %#v", bodyData)
	}
}
