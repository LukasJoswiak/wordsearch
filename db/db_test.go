package db_test

import (
    "fmt"
    "os"
    "testing"

    _ "github.com/go-sql-driver/mysql"

    "github.com/LukasJoswiak/wordsearch/app"
    "github.com/LukasJoswiak/wordsearch/config"
    "github.com/LukasJoswiak/wordsearch/db"
    "github.com/LukasJoswiak/wordsearch/models"
)

var testApp app.App

var testPuzzle = &models.Puzzle{
    ID: 1,
    URL: "123",
    Width: 3,
    Height: 2,
    Data: "abcdef",
}

var wordsAdded = &models.Words{
    PuzzleID: testPuzzle.ID,
    Words: []models.Word{
        models.Word{
            ID: 0,
            Word: "hello",
        },
        models.Word{
            ID: 0,
            Word: "world",
        },
        models.Word{
            ID: 0,
            Word: "goodbye",
        },
    },
}

func TestMain(m *testing.M) {
    setupDatabase()
    result := m.Run()
    teardownDatabase()
    os.Exit(result)
}

func setupDatabase() {
    config := config.New()
    config.Database.DatabaseName = "wordsearch_test"

    db, err := db.InitDB(config.Database)
    if err != nil {
        fmt.Errorf("failed to initialize database: %v", err)
    }

    testApp = app.App{config, db}
}

func teardownDatabase() {
    testApp.Database.DropDatabase(testApp.Config.Database.DatabaseName)
    testApp.Database.Close()
}

func TestCreatePuzzle(t *testing.T) {
    err := testApp.Database.CreatePuzzle(testPuzzle)
    if err != nil {
        t.Errorf("error creating puzzle: %v", err)
    }
}

func TestGetNonExistentPuzzle(t *testing.T) {
    puzzle, err := testApp.Database.GetPuzzle("1")
    if err != nil {
        t.Errorf("expected nil error, got %v", err)
    }
    if puzzle != nil {
        t.Errorf("expected nil puzzle, got %v", puzzle)
    }
}

func TestGetPuzzle(t *testing.T) {
    puzzle, err := testApp.Database.GetPuzzle(testPuzzle.URL)
    if err != nil {
        t.Errorf("error fetching puzzle: %v", err)
    }

    if puzzle.ID != testPuzzle.ID || puzzle.Width != testPuzzle.Width ||
            puzzle.Height != testPuzzle.Height || puzzle.Data != testPuzzle.Data {
        t.Errorf("expected %v, got %v", testPuzzle, puzzle)
    }

    testPuzzle.ID = puzzle.ID
}

func TestAddWords(t *testing.T) {
    err := testApp.Database.UpdateWords(wordsAdded)
    if err != nil {
        t.Errorf("error updating words: %v", err)
    }
}

// No words should exist yet, because the relation with the puzzle hasn't been
// created yet.
func TestGetEmptyWords(t *testing.T) {
    words, err := testApp.Database.GetWords(testPuzzle.ID)
    if err != nil {
        t.Errorf("error fetching words: %v", err)
    }

    if len(words.Words) != 0 {
        t.Errorf("expected 0 words associated with puzzle %d, got %d", testPuzzle.ID, len(words.Words))
    }
}

// Associate newly created words with puzzle.
func TestUpdatePuzzleWords(t *testing.T) {
    wordIds, err := testApp.Database.GetWordIds(wordsAdded)
    if err != nil {
        t.Errorf("error fetching word IDs: %v", err)
    }

    err = testApp.Database.UpdatePuzzleWords(testPuzzle.ID, wordIds)
    if err != nil {
        t.Errorf("error updating puzzle words: %v", err)
    }
}

func TestGetWords(t *testing.T) {
    words, err := testApp.Database.GetWords(testPuzzle.ID)
    if err != nil {
        t.Errorf("error fetching words: %v", err)
    }

    if len(words.Words) != len(wordsAdded.Words) {
        t.Errorf("expected %d words, got %d words", len(wordsAdded.Words), len(words.Words))
    }

    for _, addedWord := range wordsAdded.Words {
        foundWord := false
        for _, word := range words.Words {
            if word.Word == addedWord.Word {
                foundWord = true
            }
        }

        if !foundWord {
            t.Errorf("couldn't find word %s in words for puzzle %d", addedWord.Word, testPuzzle.ID)
        }
    }
}

func TestRemoveWords(t *testing.T) {
    var wordsRemoved = &models.Words{
        PuzzleID: testPuzzle.ID,
        Words: []models.Word{
            models.Word{
                ID: 0,
                Word: "goodbye",
            },
        },
    }

    removedWordIds, err := testApp.Database.GetWordIds(wordsRemoved)
    if err != nil {
        t.Errorf("failed to fetch word IDs: %v", err)
    }

    err = testApp.Database.RemovePuzzleWords(testPuzzle.ID, removedWordIds)
    if err != nil {
        t.Errorf("failed to remove puzzle word relations: %v", err)
    }
}

func TestGetNewWords(t *testing.T) {
    words, err := testApp.Database.GetWords(testPuzzle.ID)
    if err != nil {
        t.Errorf("error fetching words: %v", err)
    }

    if len(words.Words) != len(wordsAdded.Words) - 1 {
        t.Errorf("expected %d words associated with puzzle %d, got %d", len(wordsAdded.Words) - 1, testPuzzle.ID, len(words.Words))
    }
}
