package api

import (
    "github.com/gorilla/mux"

    "wordsearch/app"
)

type API struct {
    app *app.App
    Config *Config
}

func New(a *app.App) (api *API, err error) {
    api = &API{app: a}
    api.Config, err = InitConfig()
    if err != nil {
        return nil, err
    }
    return api, nil
}

func (api *API) Init(r *mux.Router) {
    r.HandleFunc("/puzzle/{id:[0-9]+}", api.getPuzzleHandler).Methods("GET")
}
