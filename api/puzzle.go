package api

import (
    "fmt"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

func (api *API) getPuzzleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    puzzles, err := api.app.GetPuzzle(vars["id"])
    puzzleListBytes, err := json.Marshal(puzzles)
    if err != nil {
        fmt.Println(fmt.Errorf("error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Write(puzzleListBytes)
}
