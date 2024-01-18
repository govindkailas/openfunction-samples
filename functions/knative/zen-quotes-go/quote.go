package quote

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OpenFunction/functions-framework-go/functions"
)

func init() {
	functions.HTTP("quote", ZenQuote,
		functions.WithFunctionPath("/quote"))
}

// respond back to the client with the quote author name and content.
func respond(w http.ResponseWriter, quote string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Write([]byte(quote))
}

func ZenQuote(w http.ResponseWriter, r *http.Request) {
	//get the quote from https://api.quotable.io/quotes/random .
	// this is the json response from the API [{"_id":"n4cqmf135I","author":"Alan Watts","content":"Things are as they are. Looking out into it the universe at night, we make no comparisons between right and wrong stars, nor between well and badly arranged constellations.","tags":["Philosophy","Self"],"authorSlug":"alan-watts","length":172,"dateAdded":"2022-03-12","dateModified":"2023-04-14"}]
	// we need only author and content.
	// we need to convert it to the following format
	// {"author":"<NAME>","content":"Things are as they are. Looking out into it the universe at night, we make no comparisons between
	// right and wrong stars, nor between well and badly arranged constellations."}

	var quote struct {
		Author       string   `json:"author"`
		Content      string   `json:"content"`
		Tags         []string `json:"tags"`
		AuthorSlug   string   `json:"authorSlug"`
		Length       int      `json:"length"`
		DateAdded    string   `json:"dateAdded"`
		DateModified string   `json:"dateModified"`
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://api.quotable.io/random")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		fmt.Println(err)
	}
	respond(w, quote.Author+" says: "+quote.Content)

}
