package db_test

import (
    "database/sql"
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
    if err != sql.ErrNoRows {
        t.Errorf("expected ErrNoRows, got %v", err)
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
}
