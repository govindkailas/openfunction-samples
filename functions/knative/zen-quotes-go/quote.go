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

func ZenQuote(w http.ResponseWriter, r *http.Request) {
	// skip ssl verify, somehow Go is not respecting the LetsEncrypt certificate of zenquotes.io

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://zenquotes.io/api/random")
	if err != nil {
		fmt.Println("Failed to fetch quote:", err)
		return
	}
	defer resp.Body.Close()

	//output format in the resp is like this: [ {"q":"The clock indicates the moment...but what does eternity indicate?","a":"Walt Whitman","h":"<blockquote>&ldquo;The clock indicates the moment...but what does eternity indicate?&rdquo; &mdash; <footer>Walt Whitman</footer></blockquote>"} ]
	var quote struct {
		Quote  string `json:"q"`
		Author string `json:"a"`
	}
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		fmt.Println("Failed to fetch quote:", err)
		return
	}
	fmt.Println(quote.Quote)
	fmt.Println("  - ", quote.Author)

	// response to have quote and quote author
	response := map[string]string{
		"quote":  quote.Quote,
		"author": quote.Author,
	}

	responseBytes, _ := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.Write(responseBytes)
}
