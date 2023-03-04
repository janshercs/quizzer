package quizzer

import (
	"fmt"
	"net/http"

	"github.com/janshercs/quizzer/pkg/sheets"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, sheets.GetSheetData())
}

func StartServer() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
