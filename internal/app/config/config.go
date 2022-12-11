package config

type StorageType string

const (
	StorageTypeMemory   StorageType = "memory"
	StorageTypeFile     StorageType = "file"
	StorageTypeDatabase StorageType = "database"
)

type Config struct {
	ServerAddr string `json:"server_addr" yaml:"server_addr" env:"SERVER_ADDRESS"`
	Storage    struct {
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
	c.ServerAddr = "localhost:8080"
	c.Storage.Type = StorageTypeMemory
	c.ShortURL.TTL = 0
}
