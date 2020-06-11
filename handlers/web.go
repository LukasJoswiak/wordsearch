package handlers

import (
    "html/template"
    "net/http"

    "github.com/gorilla/mux"
)

var templates = map[string]*template.Template{
    "home": template.Must(template.ParseFiles("templates/home.html", "templates/base.html")),
    "puzzle": template.Must(template.ParseFiles("templates/puzzle.html", "templates/base.html")),
}

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {// puzzle *models.Puzzle, words *models.Words) {
    err := templates[tmpl].ExecuteTemplate(w, "base", data)

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

    words, err := env.app.GetWords(puzzle.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    renderTemplate(w, "puzzle", map[string]interface{}{
        "puzzle": puzzle,
        "words": words,
    })
}
