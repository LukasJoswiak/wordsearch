package db

import (
    "database/sql"
    "log"
)

type Database struct {
    db *sql.DB
}

func InitDB(config *Config) (*Database, error) {
    // TODO: Move database password to configuration
    db, err := sql.Open("mysql", "root:password@/" + config.databaseName)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    database := &Database{db}
    database.CreateTables()

    return database, nil
}

func (db *Database) CreateTables() {
    createPuzzles := `
        CREATE TABLE IF NOT EXISTS puzzles (
            id INT NOT NULL AUTO_INCREMENT,
            url VARCHAR(32) NOT NULL,
            width SMALLINT NOT NULL,
            height SMALLINT NOT NULL,
            data TEXT,
            type TINYINT NOT NULL,
            datetime DATETIME NOT NULL,
            PRIMARY KEY (id)
        );
    `

    createWords := `
        CREATE TABLE IF NOT EXISTS words (
            id INT NOT NULL AUTO_INCREMENT,
            puzzle_id INT NOT NULL,
            word VARCHAR(255) NOT NULL,
            deleted BOOLEAN NOT NULL,
            PRIMARY KEY (id),
            FOREIGN KEY (puzzle_id) REFERENCES puzzles(id)
        );
    `

    _, err := db.db.Exec(createPuzzles)
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.db.Exec(createWords)
    if err != nil {
        log.Fatal(err)
    }
}

func (db *Database) Close() error {
    return db.db.Close()
}
