package api

import (
	"fmt"
	"net/http"
)

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello everyone!")
}
