package db

type Config struct {
    DatabaseName string
}

func InitConfig() (*Config, error) {
    config := &Config{
        DatabaseName: "wordsearch",
    }

    return config, nil
}
