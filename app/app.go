package app

import (
    "wordsearch/db"
)

type App struct {
    Database *db.Database
}

func New() (*App, error) {
    app := &App{}

    dbConfig, err := db.InitConfig()
    if err != nil {
        return nil, err
    }

    app.Database, err = db.InitDB(dbConfig)
    if err != nil {
        return nil, err
    }

    return app, nil
}

func (app *App) Close() error {
    return app.Database.Close()
}
