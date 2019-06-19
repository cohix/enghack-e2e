package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cohix/simplcrypto"
	"github.com/enghack-e2e/server/secret"
)

var enableAuth = true

var authKey string

func main() {
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/api/v1/message", auth(handleMessageRequest))

	fmt.Println("Running server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func init() {
	authKey = secret.Generate()
	fmt.Println("export ENGHACKAUTHKEY=" + authKey)
}

func handleMessageRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetMessageRequest(w, r)
	case http.MethodPost:
		handlePostMessageRequest(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleGetMessageRequest(w http.ResponseWriter, r *http.Request) {
	message, err := ioutil.ReadFile("./data")
	if err != nil {
		fmt.Printf("handleGetMessageRequest failed to read file %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Data sent!")

	w.Write(message)
}

func handlePostMessageRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(http.MaxBytesReader(w, r.Body, 1024))
	if err != nil {
		fmt.Printf("handlePostMessageRequest read body %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := ioutil.WriteFile("./data", body, 0644); err != nil {
		fmt.Printf("handlePostMessageRequest failed to write file %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Data saved!")

	w.WriteHeader(http.StatusOK)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func auth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if enableAuth {
			hmacHeader := r.Header.Get("AUTHORIZATION")
			hmac := simplcrypto.Base64URLEncode(simplcrypto.HMACWithSecretAndData(authKey, "EngHack2019!"))

			if hmacHeader != hmac {
				fmt.Printf("auth failed\n")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}
