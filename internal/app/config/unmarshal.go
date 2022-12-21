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

func Unmarshal(filename string, c *Config) error {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext = strings.TrimLeft(ext, "."); ext {
	case "yaml", "yml":
		return UnmarshalYAML(filename, c)
	case "json":
		return UnmarshalJSON(filename, c)
	}
	return errors.New("unknown config file extension")
}

func UnmarshalJSON(filename string, c *Config) error {
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

func UnmarshalYAML(filename string, c *Config) error {
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
