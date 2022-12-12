package config

type StorageType string

const (
    StorageTypeMemory   StorageType = "memory"
    StorageTypeFile     StorageType = "file"
    StorageTypeDatabase StorageType = "database"
)

type Config struct {
    Server struct {
        Addr    string `json:"addr" yaml:"addr" env:"SERVER_ADDRESS"`
        BaseURL string `json:"base_url" yaml:"base_url" env:"BASE_URL"`
    } `json:"server" yaml:"server" env:"-"`
    Storage struct {
        Type StorageType `json:"type" yaml:"type" env:"-"`
        DSN  string      `json:"dsn" yaml:"dsn" env:"-"`
        Path string      `json:"path" yaml:"path" env:"FILE_STORAGE_PATH"`
    } `json:"storage" yaml:"storage" env:"-"`
    ShortURL struct {
        TTL int `yaml:"ttl" env:"-"`
    } `json:"short_url" yaml:"short_url" env:"-"`
}

func New() *Config {
    c := Config{}
    c.SetDefaultValues()
    return &c
}

func (c *Config) SetDefaultValues() {
    c.Server.Addr = "127.0.0.1:8080"
    c.Server.BaseURL = "localhost:8080"

    c.Storage.Type = StorageTypeMemory
    c.ShortURL.TTL = 0
}
