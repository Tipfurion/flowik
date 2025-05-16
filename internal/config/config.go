//TODO: добавить валидацию конфига

package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	once     sync.Once
	Instance *Config
)

type Config struct {
	Db *struct {
		Type string `yaml:"type"`

		Pg *struct {
			ConnectionString string `yaml:"connectionString"`
		} `yaml:"pg"`

		Sqlite *struct {
			ConnectionString string `yaml:"connectionString"`
		} `yaml:"sqlite"`
	} `yaml:"db"`
}

func load() {
	var loadErr error
	path := "../../config.yaml"

	once.Do(func() {
		bytes, err := os.ReadFile(path)
		if err != nil {
			loadErr = fmt.Errorf("failed to read config file: %w", err)
		}

		var config Config

		if err := yaml.Unmarshal(bytes, &config); err != nil {
			loadErr = fmt.Errorf("failed to parse config: %w", err)
		}

		Instance = &config

	})

	if loadErr != nil {
		panic(fmt.Errorf("failed to load config: %w", loadErr))
	}

}

func Get() *Config {
	if Instance == nil {
		load()
	}
	return Instance
}
