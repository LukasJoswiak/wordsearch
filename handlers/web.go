package handlers

import (
    "fmt"
    "net/http"
)

func (env *Environment) homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Home")
}
