package app

import (
    "wordsearch/models"
)

func (app *App) GetPuzzle(url string) (*models.Puzzle, error) {
    puzzle, err := app.Database.GetPuzzle(url)
    if err != nil {
        return nil, err
    }
    return puzzle, nil
}
