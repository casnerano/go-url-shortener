// Package for application configuration
package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

// DefaultConfigFileName - default configuration file path.
const DefaultConfigFileName = "./configs/application.yaml"

// StorageType defines the type of storage
type StorageType string

// List of storage types.
const (
	StorageTypeMemory   StorageType = "memory"
	StorageTypeFile     StorageType = "file"
	StorageTypeDatabase StorageType = "database"
)

// Config is a structure with contains application configurations.
type Config struct {
	App struct {
		Secret string `json:"secret" yaml:"secret"`
	} `json:"app" yaml:"app"`
	Server struct {
		EnableHTTPS bool   `json:"enable_https" yaml:"enable_https" env:"ENABLE_HTTPS"`
		Addr        string `json:"addr" yaml:"addr" env:"SERVER_ADDRESS"`
		BaseURL     string `json:"base_url" yaml:"base_url" env:"BASE_URL"`
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

// New configuration constructor.
func New() *Config {
	return &Config{}
}

// Init from other
func (c *Config) Init() {
	c.SetDefaultValues()

	flags := parseFlags(c)

	filename := DefaultConfigFileName
	if flags.ConfigName != "" {
		filename = flags.ConfigName
	}

	if err := c.setConfigFileValues(filename); err != nil {
		log.Fatal(err.Error())
	}

	if err := c.setAppFlagValues(flags); err != nil {
		log.Fatal(err.Error())
	}

	if err := c.setEnvironmentValues(); err != nil {
		log.Fatal(err.Error())
	}

	if c.Storage.DSN != "" {
		c.Storage.Type = StorageTypeDatabase
	}

	if c.Storage.Path != "" {
		c.Storage.Type = StorageTypeFile
	}
}

// SetDefaultValues sets default values.
func (c *Config) SetDefaultValues() {
	c.App.Secret = "cfcd208495d565ef66e7dff9f98764da"

	c.Server.Addr = "127.0.0.1:8080"
	c.Server.BaseURL = "http://localhost:8080"

	c.Storage.Type = StorageTypeMemory
	c.ShortURL.TTL = 0
}

// setConfigFileValues sets values from file.
func (c *Config) setConfigFileValues(filename string) error {
	return unmarshal(filename, c)
}

// setEnvironmentValues sets values from environment variables.
func (c *Config) setEnvironmentValues() error {
	return env.Parse(c)
}

// setAppFlagValues sets values from flags.
func (c *Config) setAppFlagValues(flags *Flags) error {
	c.Server.Addr = flags.Server.Addr
	c.Server.EnableHTTPS = flags.Server.EnableHTTPS
	c.Server.BaseURL = flags.Server.BaseURL
	c.Storage.Path = flags.Storage.Path
	c.Storage.DSN = flags.Storage.DSN

	return nil
}
