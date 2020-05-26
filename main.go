package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"

    "wordsearch/app"
    "wordsearch/api"
)

func setupRoutes(api *api.API) *mux.Router {
    r := mux.NewRouter()

    api.Init(r.PathPrefix("/api").Subrouter())

    // r.HandleFunc("/", HomeHandler)
    // r.HandleFunc("/p/{id:[0-9]+}", EditHandler)
    // r.HandleFunc("/v/{id:[0-9]+}", ViewHandler)

    staticFileDirectory := http.Dir("./assets/")
    staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
    r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

    return r
}

func main() {
    app, err := app.New()
    if err != nil {
        log.Fatal(err)
    }
    defer app.Close()

    api, err := api.New(app)
    if err != nil {
        log.Fatal(err)
    }

    r := setupRoutes(api)

    server := &http.Server{
        Handler: r,
        Addr: fmt.Sprintf(":%d", api.Config.Port),
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    log.Printf("listening on port %d", 8080)
    log.Fatal(server.ListenAndServe())
}
