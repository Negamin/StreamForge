package config

import (
    "os"
    "gopkg.in/yaml.v3"
)

type Config struct {
    Pipeline struct {
        Source         string   `yaml:"source"`
        Transformations []string `yaml:"transformations"`
        Sink           string   `yaml:"sink"`
    } `yaml:"pipeline"`
}

func LoadConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var cfg Config
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}