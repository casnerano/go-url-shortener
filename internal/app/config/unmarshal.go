package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const defaultENVFilename = ".env"

// unmarshal configuration file.
func unmarshal(filename string, c *Config) error {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext = strings.TrimLeft(ext, "."); ext {
	case "yaml", "yml":
		return unmarshalYAML(filename, c)
	case "json":
		return unmarshalJSON(filename, c)
	}
	return errors.New("unknown config file extension")
}

// unmarshalJSON configuration file.
func unmarshalJSON(filename string, c *Config) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, c)
	if err != nil {
		return err
	}

	return nil
}

// unmarshalYAML configuration file.
func unmarshalYAML(filename string, c *Config) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		return err
	}

	return nil
}
