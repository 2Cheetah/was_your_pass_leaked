package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	pingHandler(w, req)
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

func TestIsLeakedHanlder(t *testing.T) {
	reqBody := []byte("password")
	req := httptest.NewRequest(http.MethodPost, "/isLeaked", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	isLeakedHandler(w, req)
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

	var expected = map[string]bool{"isLeaked": true}
	var bodyData map[string]bool
	if err := json.Unmarshal(body, &bodyData); err != nil {
		t.Errorf("Couldn't unmarshall response body: %#v", err)
	}
	if bodyData["isLeaked"] != expected["isLeaked"] {
		t.Errorf("Reponse body data is wrong: %#v", bodyData)
	}
}

// Table-driven tests for sha1HashFromString_ValidPassword
func TestSha1HashFromString(t *testing.T) {
	testCases := []struct {
		password string
		expected string
	}{
		{"password", "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8"},
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"p@ssw0rd! ", "113b8d0b7b8a30eb5380e120b696dc58c3f4852d"},
	}

	for i, testCase := range testCases {
		t.Logf("Running test case %d", i+1)
		result := sha1HashFromString(testCase.password)
		if result != testCase.expected {
			t.Errorf("sha-1 sum didn't match expected for password %s",
				testCase.password)
		}
	}
}
