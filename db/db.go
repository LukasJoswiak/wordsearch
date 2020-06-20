package db

import (
    "database/sql"
    "log"
    "os"
)

type DBError struct {
    QueryString string
    Err error
}

func (dbe DBError) Error() string {
    if dbe.Err != nil {
        return dbe.Err.Error()
    }
    return ""
}

func (dbe DBError) Query() string {
    return dbe.QueryString
}

type Database struct {
    db *sql.DB
}

func InitDB(config *Config) (*Database, error) {
    dbPassword := os.Getenv("DB_PASSWORD")
    db, err := sql.Open("mysql", "root:" + dbPassword + "@/")
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    database := &Database{db}
    database.CreateDatabase(config.DatabaseName)
    database.UseDatabase(config.DatabaseName)
    database.CreateTables()

    return database, nil
}

func (db *Database) CreateDatabase(database string) {
    createDatabase := "CREATE DATABASE IF NOT EXISTS " + database + " CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci"

    _, err := db.db.Exec(createDatabase)
    if err != nil {
        log.Fatal(err)
    }
}

func (db *Database) UseDatabase(database string) {
    useDatabase := "USE " + database

    _, err := db.db.Exec(useDatabase)
    if err != nil {
        log.Fatal(err)
    }
}

func (db *Database) DropDatabase(database string) {
    dropDatabase := "DROP DATABASE " + database

    _, err := db.db.Exec(dropDatabase)
    if err != nil {
        log.Fatal(err)
    }
}

func (db *Database) CreateTables() {
    createPuzzles := `
        CREATE TABLE IF NOT EXISTS puzzles (
            id INT NOT NULL AUTO_INCREMENT,
            url VARCHAR(32) NOT NULL,
            view_url VARCHAR(32) NOT NULL,
            width SMALLINT NOT NULL,
            height SMALLINT NOT NULL,
            data TEXT,
            type TINYINT NOT NULL,
            datetime DATETIME NOT NULL,
            PRIMARY KEY (id)
        )
    `

    createWords := `
        CREATE TABLE IF NOT EXISTS words (
            id INT NOT NULL AUTO_INCREMENT,
            word VARCHAR(255) NOT NULL UNIQUE,
            PRIMARY KEY (id)
        )
    `

    createPuzzleWords := `
        CREATE TABLE IF NOT EXISTS puzzle_words (
            word_id INT,
            puzzle_id INT,
            PRIMARY KEY (word_id, puzzle_id),
            FOREIGN KEY (word_id) REFERENCES words(id),
            FOREIGN KEY (puzzle_id) REFERENCES puzzles(id)
        )
    `

    _, err := db.db.Exec(createPuzzles)
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.db.Exec(createWords)
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.db.Exec(createPuzzleWords)
    if err != nil {
        log.Fatal(err)
    }
}

func (db *Database) Close() error {
    return db.db.Close()
}
