package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"slices"
	"strings"
)

func main() {
	startServer()
}

func startServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", pingHandler)
	mux.HandleFunc("POST /isLeaked", isLeakedHandler)

	slog.Info("Starting server")
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
	password := r.FormValue("password")
	isLeaked := checkLeaked(password)

	respMap := map[string]bool{"isLeaked": isLeaked}
	respJson, err := json.Marshal(respMap)
	if err != nil {
		slog.Warn(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

// sha1HashFromString takes a password as input and returns its SHA-1 hash as a string.
//
// Example Usage:
//
//	password := "password"
//	hashedPassword := sha1HashFromString(password)
//	fmt.Println(hashedPassword)
//	// Output: 5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8
//
// Inputs:
//
//	password (string): The password to be hashed.
//
// Outputs:
//
//	hashedPassword (string): The SHA-1 hash of the input password.
func sha1HashFromString(password string) string {
	hash := sha1.Sum([]byte(password))
	return fmt.Sprintf("%x", hash)
}

func listLeakedPasswords(passwordHash string) []string {
	firstFiveSymbols := passwordHash[:5]
	url := fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", firstFiveSymbols)
	fmt.Printf("Calling: %s\n", url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while calling api:", err.Error())
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Printf("Error while reading response body: %#v", err)
	}

	bodyStr := string(body)
	var listOfStrings []string = strings.Split(bodyStr, "\n")
	var listOfHashes []string

	for _, hash := range listOfStrings {
		hash = strings.TrimRight(hash, "\r\n")
		hash = strings.Split(hash, ":")[0]
		listOfHashes = append(listOfHashes, hash)
		// fmt.Printf("%#v", hash)
	}

	return listOfHashes
}

func checkLeaked(password string) bool {
	passwordHash := sha1HashFromString(password)
	passwordHash = strings.ToUpper(passwordHash)
	listOfHashes := listLeakedPasswords(passwordHash)
	containsHash := slices.Contains(listOfHashes, passwordHash[5:])
	return containsHash
}
