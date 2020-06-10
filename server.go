package main

import (
    "fmt"
	"net/http"
	"os"
)

func EchoServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Headers:\n")
	fmt.Fprint(w, "================================\n\n")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
	fmt.Fprint(w, "\n\n")
	fmt.Fprint(w, "Environment Variables:\n")
	fmt.Fprint(w, "================================\n\n")
	for _, e := range os.Environ() {
		fmt.Fprintf(w, "%s\n", e)
    }
	fmt.Fprint(w, "\n")
}

func main() {
    http.HandleFunc("/", EchoServer)
    http.ListenAndServe(":8080", nil)
}

