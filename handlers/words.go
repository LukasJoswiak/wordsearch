package handlers

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/schema"

    "github.com/LukasJoswiak/wordsearch/models"
)

var decoder = schema.NewDecoder()

func (env *Environment) wordsHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    url := vars["url"]

    err := r.ParseForm()
    if err != nil {
        fmt.Println(fmt.Errorf("error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    var words models.WordsForm
    err = decoder.Decode(&words, r.PostForm)
    if err != nil {
        fmt.Println(fmt.Errorf("error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    err = env.app.UpdateWords(url, words)
    if err != nil {
        fmt.Println(fmt.Errorf("error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/p/" + url, http.StatusFound)
}
