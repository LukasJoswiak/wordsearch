package handlers

import (
    "net/http"

    "github.com/gorilla/mux"

    "wordsearch/app"
)

type Environment struct {
    app *app.App
}

func New(app *app.App) *Environment {
    env := &Environment{app: app}
    return env
}

func (env *Environment) Init(r *mux.Router) {
    staticFileDirectory := http.Dir("./assets/")
    staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
    r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

    r.HandleFunc("/", env.homeHandler)
    r.HandleFunc("/p/{url:[0-9]+}", env.editHandler)
    // r.HandleFunc("/v/{url:[0-9]+}", env.viewHandler)

    s := r.PathPrefix("/api").Subrouter()
    s.HandleFunc("/puzzle/{url:[0-9]+}", env.getPuzzleHandler).Methods("GET")
    s.HandleFunc("/puzzle/create", env.createPuzzleHandler).Methods("POST");

    s.HandleFunc("/words/update/{url:[0-9]+}", env.wordsHandler).Methods("POST")
}
