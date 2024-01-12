package function

import (
	"context"
	"fmt"
	"net/http"

	"rsc.io/quote"
)

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {

	quote_string := quote.Go()

	fmt.Println("Received request at /")
	fmt.Println("Sending quote", quote_string)
	// Write the quote to the client
	//res.WriteHeader(http.StatusOK)
	//res.Write([]byte(quote_string))

	fmt.Fprint(res, quote_string) // echo to quote to caller
}
