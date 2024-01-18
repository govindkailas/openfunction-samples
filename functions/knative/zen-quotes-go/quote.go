package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/quote", quoteHandler)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request, going to call the Quote API")
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
		Quote string `json:"content"`
	}

	err = json.Unmarshal(body, &quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "\"%s\" - %s", quote.Content, quote.Author )
	//print only quote.content, quote.author as json to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
	// json.NewEncoder(w).Encode(map[string]string{"quote": quote.Content, "author": quote.Author})

	log.Println("Response sent")
}
