package handlers

import (
    "fmt"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

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
