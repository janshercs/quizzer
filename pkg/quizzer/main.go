package quizzer

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Jans!")
}

func StartServer() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
