package quizzer

import (
	"log"
	"net/http"
	"text/template"

	"github.com/janshercs/quizzer/pkg/sheets"
)

const (
	formHTML = "./static/form.html"
	quizHTML = "./static/quiz.html"
)

func StartServer() {
	http.HandleFunc("/", handleForm)
	log.Printf("Serving on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sheetUrl := r.FormValue("sheetUrl")

		data := sheets.GetSheetData(sheetUrl)

		quizTemplate := template.Must(template.ParseFiles(quizHTML))
		err := quizTemplate.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Render the form page using a template
		formTemplate := template.Must(template.ParseFiles(formHTML))
		err := formTemplate.Execute(w, "hi")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
