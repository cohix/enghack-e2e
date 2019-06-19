package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cohix/enghack-e2e/server/auth"
)

func main() {
	http.HandleFunc("/api/v1/message", auth.Verify(handleMessageRequest))

	fmt.Println("Running server on port 8080")
	http.ListenAndServe(":8080", nil)
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
