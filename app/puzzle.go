package app

import (
    "math/rand"
    "strconv"
    "strings"
    "regexp"
    "time"

    "github.com/LukasJoswiak/wordsearch/models"
)

const (
    min = 1000000
    max = 9999999
)

func (app *App) GetPuzzle(url string) (*models.Puzzle, error) {
    puzzle, err := app.Database.GetPuzzle(url)
    if err != nil {
        return nil, err
    }
    puzzle.URL = url
    puzzle.Data = strings.Replace(puzzle.Data, ",", "\n", -1)

    return puzzle, nil
}

func (app *App) CreatePuzzle(body string) (string, error) {
    rand.Seed(time.Now().UnixNano())
    url := strconv.Itoa(rand.Intn(max - min) + min)

    re := regexp.MustCompile(`\r?\n`)
    body = re.ReplaceAllString(body, ",")

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

func (app *App) UpdatePuzzle(url string, body string) error {
    puzzle, err := app.Database.GetPuzzle(url)
    if err != nil {
        return err
    }
    puzzle.Data = body

    err = app.Database.UpdatePuzzle(puzzle)
    if err != nil {
        return err
    }

    return nil
}
