package handlers

import (
    "html/template"
    "net/http"

    "github.com/gorilla/mux"

    "wordsearch/models"
)

var templates = template.Must(template.ParseFiles("templates/home.html", "templates/puzzle.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, puzzle *models.Puzzle) {
    err := templates.ExecuteTemplate(w, tmpl + ".html", puzzle)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func (env *Environment) homeHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "home", nil)
}

func (env *Environment) editHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    url := vars["url"]
    puzzle, err := env.app.GetPuzzle(url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    renderTemplate(w, "puzzle", puzzle)
}
