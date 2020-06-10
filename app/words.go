package app

import (
    "wordsearch/models"
)

func (app *App) GetWords(puzzleId int) (*models.Words, error) {
    words, err := app.Database.GetWords(puzzleId)
    if err != nil {
        return nil, err
    }
    return words, nil
}

func (app *App) UpdateWords(url string, words models.WordsForm) error {
    // Get info about the puzzle such as the puzzle ID.
    puzzle, err := app.Database.GetPuzzle(url)
    if err != nil {
        return err
    }

    // Translate form into a set of words added and words removed.
    wordsAdded := &models.Words{PuzzleID: puzzle.ID}
    wordsRemoved := &models.Words{PuzzleID: puzzle.ID}

    for _, word := range words.Words {
        wordString := word.Word
        existingWordString := word.ExistingWord

        // Ignore words that were unchanged.
        if wordString == existingWordString {
            continue
        }

        wordsRemoved.Words = append(wordsRemoved.Words, models.Word{ID: 0, Word: existingWordString})
        if len(wordString) > 0 {
            wordsAdded.Words = append(wordsAdded.Words, models.Word{ID: 0, Word: wordString})
        }
    }

    // Get word IDs of deleted words.
    removedWordIds, err := app.Database.GetWordIds(wordsRemoved)
    if err != nil {
        return err
    }

    // Remove relations between removed words and the puzzle.
    err = app.Database.RemovePuzzleWords(puzzle.ID, removedWordIds)
    if err != nil {
        return err
    }

    // Add any new words to the database of all words.
    err = app.Database.UpdateWords(wordsAdded)
    if err != nil {
        return err
    }

    // Get word IDs of newly added words.
    wordIds, err := app.Database.GetWordIds(wordsAdded)
    if err != nil {
        return err
    }

    // Create a relation between the puzzle and any new words.
    err = app.Database.UpdatePuzzleWords(puzzle.ID, wordIds)
    if err != nil {
        return err
    }

    return nil
}
