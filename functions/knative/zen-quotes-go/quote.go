package quote

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/OpenFunction/functions-framework-go/functions"
)

func init() {
	functions.HTTP("quote", ZenQuote,
		functions.WithFunctionPath("/quote"))
}

type Quote struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

func ZenQuote(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		fmt.Println("Failed to fetch quote:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		fmt.Println("Failed to parse response body:", err)
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
