package config

type StorageType string

const (
	StorageTypeMemory   StorageType = "memory"
	StorageTypeDatabase StorageType = "database"
)

type Config struct {
	ServerAddr string `json:"server_addr" yaml:"server_addr"`
	Storage    struct {
		Type StorageType `json:"type" yaml:"type"`
		DSN  string      `json:"dsn" yaml:"dsn"`
	} `json:"storage" yaml:"storage"`
	ShortURL struct {
		TTL int `yaml:"ttl"`
	} `json:"short_url" yaml:"short_url"`
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
