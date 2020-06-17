package handlers

import (
    "html/template"
    "net/http"

    "github.com/gorilla/mux"
)

var templates = map[string]*template.Template{
    "home": template.Must(template.ParseFiles("templates/home.html", "templates/base.html")),
    "puzzle": template.Must(template.ParseFiles("templates/puzzle.html", "templates/base.html")),
    "view_puzzle": template.Must(template.ParseFiles("templates/view_puzzle.html", "templates/base.html")),
    "edit_puzzle": template.Must(template.ParseFiles("templates/edit_puzzle.html", "templates/base.html")),
    "error": template.Must(template.ParseFiles("templates/error.html", "templates/base.html")),
}

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) error {
    err := templates[tmpl].ExecuteTemplate(w, "base", data)
    if err != nil {
        return err
    }
    return nil
}

func (env *Environment) homeHandler(w http.ResponseWriter, r *http.Request) error {
    err := renderTemplate(w, "home", nil)
    if err != nil {
        return StatusError{500, err}
    }
    return nil
}

func (env *Environment) editWordsHandler(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    url := vars["url"]

    puzzle, err := env.app.GetPuzzle(url)
    if err != nil {
        return StatusError{500, err}
    }

    if puzzle == nil {
        return StatusError{404, nil}
    }

    words, err := env.app.GetWords(puzzle.ID)
    if err != nil {
        return StatusError{500, err}
    }

    solvedPuzzle := env.app.SolvePuzzle(puzzle, words)

    err = renderTemplate(w, "puzzle", map[string]interface{}{
        "solvedPuzzle": solvedPuzzle,
        "words": words,
    })
    if err != nil {
        return StatusError{500, err}
    }

    return nil
}

func (env *Environment) viewPuzzleHandler(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    url := vars["url"]

    puzzle, err := env.app.GetPuzzleByViewUrl(url)
    if err != nil {
        return StatusError{500, err}
    }

    if puzzle == nil {
        return StatusError{404, nil}
    }

    words, err := env.app.GetWords(puzzle.ID)
    if err != nil {
        return StatusError{500, err}
    }

    solvedPuzzle := env.app.SolvePuzzle(puzzle, words)

    err = renderTemplate(w, "view_puzzle", map[string]interface{}{
        "solvedPuzzle": solvedPuzzle,
        "words": words,
    })
    if err != nil {
        return StatusError{500, err}
    }

    return nil
}

func (env *Environment) editPuzzleHandler(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    url := vars["url"]

    puzzle, err := env.app.GetFormattedPuzzle(url)
    if err != nil {
        return StatusError{500, err}
    }

    err = renderTemplate(w, "edit_puzzle", map[string]interface{}{
        "puzzle": puzzle,
    })
    if err != nil {
        return StatusError{500, err}
    }

    return nil
}

func (env *Environment) catchAllHandler(w http.ResponseWriter, r *http.Request) error {
    return StatusError{404, nil}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    renderTemplate(w, "error", map[string]interface{}{
        "status": status,
    })
}
