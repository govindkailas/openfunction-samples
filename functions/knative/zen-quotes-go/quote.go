package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/quote", quoteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.quotable.io/random")
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

	var quote struct {
		Author  string `json:"author"`
		Content string `json:"content"`
	}

	err = json.Unmarshal(body, &quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s says: %s", quote.Author, quote.Content)
}
