package app

import (
    "wordsearch/models"
)

func (app *App) UpdateWords(url string, words models.Words) error {
    // Convert Words struct into array of word strings.
    var wordStrings []string
    for _, word := range words.Words {
        wordString := word.Word
        if len(wordString) > 0 {
            wordStrings = append(wordStrings, wordString)
        }
    }

    err := app.Database.UpdateWords(url, wordStrings)
    if err != nil {
        return err
    }

    return nil
}
