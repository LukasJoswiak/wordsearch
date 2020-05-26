package db

type Config struct {
    databaseName string
}

func InitConfig() (*Config, error) {
    config := &Config{
        databaseName: "wordsearch",
    }

    return config, nil
}
