package handlers

import (
    "context"
    // "encoding/base64"
    "fmt"
    // "io/ioutil"
    "net/http"
    // "os"

    "github.com/gorilla/mux"
    vision "cloud.google.com/go/vision/apiv1"
)

func (env *Environment) createPuzzleHandler(w http.ResponseWriter, r *http.Request) error {
    err := r.ParseForm()
    if err != nil {
        return StatusError{500, err}
    }

    body := r.PostFormValue("body")
    if len(body) == 0 {
        return StatusError{400, nil}
    }

    url, err := env.app.CreatePuzzle(body, 0)
    if err != nil {
        return StatusError{500, err}
    }

    http.Redirect(w, r, "/p/" + url, http.StatusFound)
    return nil
}

func (env *Environment) updatePuzzleHandler(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    url := vars["url"]

    err := r.ParseForm()
    if err != nil {
        return StatusError{500, err}
    }

    body := r.PostFormValue("body")
    if len(body) == 0 {
        return StatusError{400, nil}
    }

    err = env.app.UpdatePuzzle(url, body)
    if err != nil {
        return StatusError{500, err}
    }

    http.Redirect(w, r, "/p/" + url, http.StatusFound)
    return nil
}

func (env *Environment) clonePuzzleHandler(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    viewUrl := vars["view_url"]

    url, err := env.app.ClonePuzzle(viewUrl)
    if err != nil {
        return StatusError{500, err}
    }

    http.Redirect(w, r, "/p/" + url, http.StatusFound)
    return nil
}

func (env *Environment) uploadPuzzleHandler(w http.ResponseWriter, r *http.Request) error {
    // TODO: Test this with different values
    // Store 10MB of form data in memory, rest on disk.
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        return StatusError{500, err}
    }

    file, handler, err := r.FormFile("puzzle")
    if err != nil {
        return StatusError{500, err}
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Read uploaded image contents into byte array.
    /*
    contents, err := ioutil.ReadAll(file)
    if err != nil {
        return StatusError{500, err}
    }

    // TODO: Can read first 512 bytes if not using the rest
    filetype := http.DetectContentType(contents)
    if filetype != "image/jpg" && filetype != "image/jpeg" && filetype != "image/png" {
        return StatusError{415, nil}
    }
    fmt.Println("filetype: ", filetype)
    */

    // encoded := base64.StdEncoding.EncodeToString(contents)

    ctx := context.Background()

    client, err := vision.NewImageAnnotatorClient(ctx)
    if err != nil {
        return StatusError{500, err}
    }

    image, err := vision.NewImageFromReader(file)
    if err != nil {
        return StatusError{500, err}
    }
    annotations, err := client.DetectTexts(ctx, image, nil, 100)
    if err != nil {
        return StatusError{500, err}
    }

    if len(annotations) == 0 {
            fmt.Fprintln(w, "No text found.")
    } else {
        fmt.Fprintln(w, "Text:")
        for _, annotation := range annotations {
                fmt.Fprintf(w, "%q\n", annotation.Description)
        }
    }

    return nil
}
