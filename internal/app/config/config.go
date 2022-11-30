package config

type StorageType string

const (
	StorageTypeMemory   StorageType = "memory"
	StorageTypeDatabase StorageType = "database"
)

type Config struct {
	ServerAddr string `json:"serverAddr" yaml:"serverAddr"`
	Storage    struct {
		Type StorageType `json:"type" yaml:"type"`
		DSN  string      `json:"DSN" yaml:"DSN"`
	} `json:"repository" yaml:"repository"`
	ShortURL struct {
		TTL int `yaml:"TTL"`
	} `json:"shortURL" yaml:"shortURL"`
}

func New() *Config {
	c := Config{}
	c.SetDefaultValues()
	return &c
}

func (c *Config) SetDefaultValues() {
	c.ServerAddr = "localhost:8080"
	c.Storage.Type = StorageTypeMemory
	c.ShortURL.TTL = 0
}
