package db

import (
    "database/sql"
    "log"
    "os"
    "time"
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
    dbUser := os.Getenv("DB_USER")
    if len(dbUser) == 0 {
        log.Fatal("missing DB_USER environment variable")
    }
    dbPassword := os.Getenv("DB_PASSWORD")
    if len(dbPassword) == 0 {
        log.Fatal("missing DB_PASSWORD environment variable")
    }
    db, err := sql.Open("mysql", dbUser + ":" + dbPassword + "@/")
    if err != nil {
        log.Fatal(err)
    }

    CreateDatabase(db, config.DatabaseName)
    db.Close()

    // Specify database name in DSN to prevent issues with connection pools.
    db, err = sql.Open("mysql", dbUser + ":" + dbPassword + "@/" + config.DatabaseName)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(time.Minute * 5)

    database := &Database{db}
    database.CreateTables()

    return database, nil
}

func CreateDatabase(db *sql.DB, database string) {
    createDatabase := "CREATE DATABASE IF NOT EXISTS " + database + " CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"

    _, err := db.Exec(createDatabase)
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
