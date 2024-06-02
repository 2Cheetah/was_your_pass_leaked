package helper

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"slices"
	"strings"
)

func Sha1HashFromString(password string) string {
	hash := sha1.Sum([]byte(password))
	return fmt.Sprintf("%x", hash)
}

func ListLeakedPasswords(passwordHash string) []string {
	firstFiveSymbols := passwordHash[:5]
	slog.Info("Calculated password hash.", "First 5 symbols", firstFiveSymbols)
	url := fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", firstFiveSymbols)
	slog.Info("Calling HIBP API.", "URL:", url)

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

func CheckLeaked(password string) bool {
	passwordHash := Sha1HashFromString(password)
	passwordHash = strings.ToUpper(passwordHash)
	listOfHashes := ListLeakedPasswords(passwordHash)
	containsHash := slices.Contains(listOfHashes, passwordHash[5:])
	return containsHash
}
