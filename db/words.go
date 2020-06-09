package db

import (
    "fmt"
    "strings"
)

func (db *Database) UpdateWords(url string, words []string) error {
    if len(words) == 0 {
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

    for _, word := range words {
        parameters = append(parameters, parameter)
        values = append(values, word)
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
