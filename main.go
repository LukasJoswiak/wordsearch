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

    "github.com/LukasJoswiak/wordsearch/app"
    "github.com/LukasJoswiak/wordsearch/config"
    "github.com/LukasJoswiak/wordsearch/handlers"
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

    logFile, err := os.Create(config.LogFile)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()
    loggedRouter := ghandlers.LoggingHandler(logFile, r)

    // Write errors to error log.
    errorLogFile, err := os.Create(config.ErrorLogFile)
    if err != nil {
        log.Fatal(err)
    }
    defer errorLogFile.Close()
    log.SetOutput(errorLogFile)

    server := &http.Server{
        Handler: ghandlers.CompressHandler(loggedRouter),
        Addr: fmt.Sprintf(":%d", config.Port),
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    fmt.Printf("server listening on port %d\n", config.Port)
    log.Fatal(server.ListenAndServe())
}
