package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type ServerConfig struct {
	Server Config `yaml:"server"`
}

type Config struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	RedisPort int    `yaml:"redis-port"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config ServerConfig
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %v", err)
	}

	// Now decode the content
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config.Server, nil
}
