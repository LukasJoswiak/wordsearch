package app

import (
    "math/rand"
    "strconv"
    "time"

    "github.com/LukasJoswiak/wordsearch/models"
)

const (
    min = 1000000
    max = 9999999
)

func (app *App) CreatePuzzle(body string) (string, error) {
    rand.Seed(time.Now().UnixNano())
    url := strconv.Itoa(rand.Intn(max - min) + min)

    puzzle := &models.Puzzle{
        URL: url,
        Data: body,
    }
    err := app.Database.CreatePuzzle(puzzle)
    if err != nil {
        return "", err
    }

    return url, nil
}

func (app *App) GetPuzzle(url string) (*models.Puzzle, error) {
    puzzle, err := app.Database.GetPuzzle(url)
    if err != nil {
        return nil, err
    }
    puzzle.URL = url
    return puzzle, nil
}
