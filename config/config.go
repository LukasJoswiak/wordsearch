package config

import (
    "encoding/json"
    "log"
    "os"
    "path/filepath"
    "runtime"

    "github.com/LukasJoswiak/wordsearch/db"
)

type Config struct {
    Database      *db.Config `json:"database"`
    Port          int        `json:"port"`
    LogFile       string     `json:"logFile"`
}

func New() (*Config) {
    // Get path to config directory.
    _, b, _, _ := runtime.Caller(0)
    basepath   := filepath.Dir(b)

    path := filepath.Join(basepath, "config.json")
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    config := &Config{}
    err = decoder.Decode(&config)
    if err != nil {
        log.Fatal(err)
    }

    return config
}
