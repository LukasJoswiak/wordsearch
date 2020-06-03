package handlers

import (
    "fmt"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

func (env *Environment) createPuzzleHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        fmt.Println(fmt.Errorf("error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    body := r.PostFormValue("body")
    url, err := env.app.CreatePuzzle(body)
    if err != nil {
        fmt.Println(fmt.Errorf("error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/p/" + url, http.StatusFound)
}

func (env *Environment) getPuzzleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    puzzle, err := env.app.GetPuzzle(vars["url"])
    puzzleListBytes, err := json.Marshal(puzzle)
    if err != nil {
        fmt.Println(fmt.Errorf("error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Write(puzzleListBytes)
}
