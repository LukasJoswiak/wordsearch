package db

import (
    "database/sql"
    "log"
    "time"

    "wordsearch/models"
)

// TODO: Return ID?
func (db *Database) CreatePuzzle(puzzle *models.Puzzle) error {
    _, err := db.db.Exec(`INSERT INTO puzzles (url, width, height, data, type, datetime) VALUES (?, ?, ?, ?, ?, ?)`, "1234", puzzle.Width, puzzle.Height, puzzle.Data, 0, time.Now())
    return err
}

func (db *Database) GetPuzzle(url string) (*models.Puzzle, error) {
    puzzle := &models.Puzzle{}

    // TODO: Change id back to url
    row := db.db.QueryRow(`SELECT id, width, height, data FROM puzzles WHERE id = ?`, url)
    err := row.Scan(&puzzle.ID, &puzzle.Width, &puzzle.Height, &puzzle.Data)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, err
        } else {
            log.Fatal(err)
        }
    }

    return puzzle, nil
}
