package handlers

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/schema"

    "github.com/LukasJoswiak/wordsearch/models"
)

var decoder = schema.NewDecoder()

func (env *Environment) wordsHandler(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    url := vars["url"]

    err := r.ParseForm()
    if err != nil {
        return StatusError{500, err}
    }

    var words models.WordsForm
    err = decoder.Decode(&words, r.PostForm)
    if err != nil {
        return StatusError{500, err}
    }

    err = env.app.UpdateWords(url, words)
    if err != nil {
        return StatusError{500, err}
    }

    http.Redirect(w, r, "/p/" + url, http.StatusFound)
    return nil
}
