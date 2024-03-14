package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PaesslerAG/jsonpath"
)

func main() {
	http.HandleFunc("/", redirHandler)
	http.ListenAndServe(":8080", nil)
}

func redirHandler(w http.ResponseWriter, r *http.Request) {
	queryURL := r.URL.Query().Get("url")
	queryPath := r.URL.Query().Get("query")

	// Realizar a chamada GET para a URL fornecida
	resp, err := http.Get(queryURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Ler o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Decodificar o corpo da resposta JSON
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Aplicar o JSONPath para extrair a URL de download
	result, err := jsonpath.Get(queryPath, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verificar se o resultado é uma string
	downloadURL, ok := result.(string)
	if !ok {
		http.Error(w, "Result is not a string", http.StatusInternalServerError)
		return
	}

	// Redirecionar o usuário para a URL de download
	http.Redirect(w, r, downloadURL, http.StatusFound)
}
