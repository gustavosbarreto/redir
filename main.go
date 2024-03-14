package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PaesslerAG/jsonpath"
)

func main() {
	fmt.Println("Starting...")

	http.HandleFunc("/", redirHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func redirHandler(w http.ResponseWriter, r *http.Request) {
	queryURL := r.URL.Query().Get("url")
	queryPath := r.URL.Query().Get("query")

	resp, err := http.Get(queryURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := jsonpath.Get(queryPath, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	downloadURL, ok := result.(string)
	if !ok {
		http.Error(w, "Result is not a string", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, downloadURL, http.StatusFound)
}
