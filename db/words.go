package db

import (
    "fmt"
    "strings"

    "github.com/LukasJoswiak/wordsearch/models"
)

// Returns a list of words for the given puzzle ID.
func (db *Database) GetWords(puzzleId int) (*models.Words, error) {
    words := &models.Words{}

    sql := `SELECT w.word
            FROM words w
            INNER JOIN puzzle_words pw
            ON pw.word_id = w.id
            WHERE pw.puzzle_id = ?`

    rows, err := db.db.Query(sql, puzzleId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        word := models.Word{}
        err = rows.Scan(&word.Word)
        if err != nil {
            return nil, err
        }

        words.Words = append(words.Words, word)
    }

    return words, nil
}

// Returns a list of word IDs for the given words.
func (db *Database) GetWordIds(words *models.Words) ([]models.Word, error) {
    var wordIds []models.Word

    if len(words.Words) == 0 {
        return wordIds, nil
    }

    sql := `SELECT id, word
            FROM words
            WHERE word IN (%s)`

    values := []interface{}{}

    var parameters []string
    const parameter = "?"

    for _, word := range words.Words {
        parameters = append(parameters, parameter)
        values = append(values, word.Word)
    }
    sql = fmt.Sprintf(sql, strings.Join(parameters, ","))

    rows, err := db.db.Query(sql, values...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        word := models.Word{}
        err = rows.Scan(&word.ID, &word.Word)
        if err != nil {
            return nil, err
        }

        wordIds = append(wordIds, word)
    }

    return wordIds, nil
}

// Inserts the list of words into the `words` table. Words that already exist
// in the table will not be re-inserted.
func (db *Database) UpdateWords(words *models.Words) error {
    if len(words.Words) == 0 {
        return nil
    }

    // Inserts multiple rows at once while ignoring any words that already
    // exist in the words table.
    sql := `INSERT INTO words (word)
            SELECT w.word
            FROM (
                %s
            ) AS w
            WHERE NOT EXISTS (SELECT 1
                              FROM words w2
                              WHERE w2.word = w.word)`

    values := []interface{}{}

    parameters := []string{"SELECT ? AS word"}
    const parameter = "UNION ALL\nSELECT ?"

    for _, word := range words.Words {
        parameters = append(parameters, parameter)
        values = append(values, word.Word)
    }
    // First word was already inserted in parameters, so remove it.
    parameters = parameters[0: len(parameters) - 1]
    sql = fmt.Sprintf(sql, strings.Join(parameters, "\n"))

    stmt, err := db.db.Prepare(sql)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(values...)
    if err != nil {
        return err
    }

    return nil
}

// Creates a separate relation between each word ID and the puzzle with the
// given URL.
func (db *Database) UpdatePuzzleWords(puzzleId int, wordIds []models.Word) error {
    if len(wordIds) == 0 {
        return nil
    }

    sql := `INSERT INTO puzzle_words (word_id, puzzle_id)
            SELECT p.word_id, p.puzzle_id
            FROM (
                %s
            ) AS p
            WHERE NOT EXISTS (SELECT 1
                              FROM puzzle_words p2
                              WHERE p2.word_id = p.word_id
                                  AND p2.puzzle_id = p.puzzle_id)`

    values := []interface{}{}

    parameters := []string{"SELECT ? AS word_id, ? AS puzzle_id"}
    const parameter = "UNION ALL\nSELECT ?, ?"

    for _, word := range wordIds {
        parameters = append(parameters, parameter)
        values = append(values, word.ID)
        values = append(values, puzzleId)
    }
    // First word was already inserted in parameters, so remove it.
    parameters = parameters[0: len(parameters) - 1]
    sql = fmt.Sprintf(sql, strings.Join(parameters, "\n"))

    stmt, err := db.db.Prepare(sql)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(values...)
    if err != nil {
        return err
    }

    return nil
}

// Removes the relation between the puzzle with the given ID and any word with
// its ID specified in the given list.
func (db *Database) RemovePuzzleWords(puzzleId int, wordIds []models.Word) error {
    if len(wordIds) == 0 {
        return nil
    }

    sql := `DELETE FROM puzzle_words
            WHERE puzzle_id = ?
                AND word_id IN (%s)`

    values := []interface{}{puzzleId}

    var parameters []string
    const parameter = "?"

    for _, word := range wordIds {
        parameters = append(parameters, parameter)
        values = append(values, word.ID)
    }
    sql = fmt.Sprintf(sql, strings.Join(parameters, ","))

    stmt, err := db.db.Prepare(sql)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(values...)
    if err != nil {
        return err
    }

    return nil
}
