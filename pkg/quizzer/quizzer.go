package quizzer

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/janshercs/quizzer/pkg/quiz"
	"github.com/janshercs/quizzer/pkg/sheets"
)

const (
	formHTML = "./static/form.html"
	quizHTML = "./static/quiz.html"
)

func StartServer() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handleForm)

	log.Printf("Serving on: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sheetUrl := r.FormValue("sheetUrl")

		data := quiz.GetQuizFromSheetValues(sheets.GetSheetRawData(sheetUrl))
		log.Printf("%v", data)

		quizTemplate := template.Must(template.ParseFiles(quizHTML))
		err := quizTemplate.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Render the form page using a template
		formTemplate := template.Must(template.ParseFiles(formHTML))
		err := formTemplate.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
