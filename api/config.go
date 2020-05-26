package api

type Config struct {
    Port int
}

func InitConfig() (*Config, error) {
    config := &Config{
        Port: 8080,
    }
    return config, nil
}
