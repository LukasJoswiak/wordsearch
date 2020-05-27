package config

import (
    "wordsearch/db"
)

type Config struct {
    Database *db.Config
    Port int
}

// TODO: Update to read from file
func New() (*Config) {
    config := &Config{
        Database: &db.Config{
            DatabaseName: "wordsearch",
        },
        Port: 8080,
    }

    return config
}
