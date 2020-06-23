package handlers

import (
    "log"
    "net/http"
    "path/filepath"
    "runtime"

    "github.com/gorilla/mux"

    "github.com/LukasJoswiak/wordsearch/app"
)

type Error interface {
    error
    Status() int
}

type StatusError struct {
    Code int
    Err error
}

type Handler struct {
    Handle func(w http.ResponseWriter, r *http.Request) error
}

type Environment struct {
    app *app.App
}

func New(app *app.App) *Environment {
    env := &Environment{app: app}
    return env
}

func (env *Environment) Init(r *mux.Router) {
    var _, b, _, _ = runtime.Caller(0)
    var basepath = filepath.Dir(b)

    staticFileDirectory := http.Dir(filepath.Join(basepath, "../assets"))
    staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
    r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

    r.Handle("/favicon.ico", Handler{env.faviconHandler})

    r.Handle("/", Handler{env.homeHandler}).Methods("GET")
    r.Handle("/p/{url:[0-9]+}", Handler{env.solvePuzzleHandler}).Methods("GET")
    r.Handle("/p/{url:[0-9]+}/edit", Handler{env.editPuzzleHandler}).Methods("GET")
    r.Handle("/v/{url:[0-9]+}", Handler{env.viewPuzzleHandler})

    s := r.PathPrefix("/puzzle").Subrouter()
    s.Handle("/create", Handler{env.createPuzzleHandler}).Methods("POST");
    s.Handle("/{url:[0-9]+}/update", Handler{env.updatePuzzleHandler}).Methods("POST")
    s.Handle("/{view_url:[0-9]+}/clone", Handler{env.clonePuzzleHandler}).Methods("POST")

    s.Handle("/{url:[0-9]+}/words", Handler{env.wordsHandler}).Methods("POST")

    r.PathPrefix("/").Handler(Handler{env.catchAllHandler})
}

func (se StatusError) Error() string {
    if se.Err != nil {
        return se.Err.Error()
    }
    return ""
}

func (se StatusError) Status() int {
    return se.Code
}

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    err := handler.Handle(w, r)
    if err != nil {
        switch e := err.(type) {
        case Error:
            log.Printf("\"%s %s\" %d: %s\n", r.Method, r.URL, e.Status(), e)
            errorHandler(w, r, e.Status())
        default:
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
    }
}
