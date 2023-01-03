package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

const DefaultConfigFileName = "./configs/application.yaml"

type StorageType string

const (
	StorageTypeMemory   StorageType = "memory"
	StorageTypeFile     StorageType = "file"
	StorageTypeDatabase StorageType = "database"
)

type Config struct {
	App struct {
		Secret string `json:"secret" yaml:"secret"`
	} `json:"app" yaml:"app"`
	Server struct {
		Addr    string `json:"addr" yaml:"addr" env:"SERVER_ADDRESS"`
		BaseURL string `json:"base_url" yaml:"base_url" env:"BASE_URL"`
	} `json:"server" yaml:"server" env:"-"`
	Storage struct {
		Type StorageType `json:"type" yaml:"type" env:"-"`
		DSN  string      `json:"dsn" yaml:"dsn" env:"DATABASE_DSN"`
		Path string      `json:"path" yaml:"path" env:"FILE_STORAGE_PATH"`
	} `json:"storage" yaml:"storage" env:"-"`
	ShortURL struct {
		TTL int `yaml:"ttl" env:"-"`
	} `json:"short_url" yaml:"short_url" env:"-"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) SetDefaultValues() {
	c.App.Secret = "cfcd208495d565ef66e7dff9f98764da"

	c.Server.Addr = "127.0.0.1:8080"
	c.Server.BaseURL = "http://localhost:8080"

	c.Storage.Type = StorageTypeMemory
	c.ShortURL.TTL = 0
}

func (c *Config) SetConfigFileValues() error {
	return Unmarshal(DefaultConfigFileName, c)
}

func (c *Config) SetEnvironmentValues() error {
	return env.Parse(c)
}

func (c *Config) SetAppFlagValues() error {
	flag.StringVar(&c.Server.Addr, "a", c.Server.Addr, "Server addr")
	flag.StringVar(&c.Server.BaseURL, "b", c.Server.BaseURL, "Base URL")
	flag.StringVar(&c.Storage.Path, "f", c.Storage.Path, "File storage path")
	flag.StringVar(&c.Storage.DSN, "d", c.Storage.DSN, "Database connection DSN")

	flag.Parse()

	return nil
}
