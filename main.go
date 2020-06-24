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

    logFile, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()
    if _, err = logFile.WriteString("starting server\n"); err != nil {
        log.Fatal(err)
    }
    loggedRouter := ghandlers.LoggingHandler(logFile, r)

    // Write errors to error log.
    errorLogFile, err := os.OpenFile(config.ErrorLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer errorLogFile.Close()
    if _, err = errorLogFile.WriteString("starting server\n"); err != nil {
        log.Fatal(err)
    }
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
