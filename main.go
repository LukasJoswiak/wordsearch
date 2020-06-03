package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
    ghandlers "github.com/gorilla/handlers"
    _ "github.com/go-sql-driver/mysql"

    "wordsearch/app"
    "wordsearch/config"
    "wordsearch/handlers"
)

func main() {
    config := config.New()

    app, err := app.New(config)
    if err != nil {
        log.Fatal(err)
    }
    defer app.Close()

    r := mux.NewRouter()
    env := handlers.New(app)
    env.Init(r)

    // TODO: log to file
    loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)

    server := &http.Server{
        Handler: ghandlers.CompressHandler(loggedRouter),
        Addr: fmt.Sprintf(":%d", config.Port),
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    log.Printf("listening on port %d", 8080)
    log.Fatal(server.ListenAndServe())
}
