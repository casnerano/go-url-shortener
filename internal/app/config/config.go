package config

type StorageType string

const (
	STORAGE_TYPE_MEMORY   StorageType = "memory"
	STORAGE_TYPE_DATABASE StorageType = "database"
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
	c.Storage.Type = STORAGE_TYPE_MEMORY
	c.ShortURL.TTL = 60
}
