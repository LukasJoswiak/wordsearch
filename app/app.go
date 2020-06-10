package app

import (
    "github.com/LukasJoswiak/wordsearch/config"
    "github.com/LukasJoswiak/wordsearch/db"
)

type App struct {
    Config *config.Config
    Database *db.Database
}

func New(config *config.Config) (*App, error) {
    db, err := db.InitDB(config.Database)
    if err != nil {
        return nil, err
    }

    return &App{config, db}, nil
}

func (app *App) Close() error {
    return app.Database.Close()
}
