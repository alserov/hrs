package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
}

const (
	Local      = "local"
	Production = "production"
)

func MustLoad() *Config {
	path := fetchPath()

	b, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read file: " + err.Error())
	}

	var cfg Config
	if err = yaml.Unmarshal(b, &cfg); err != nil {
		panic("failed to parse configs: " + err.Error())
	}

	return &cfg
}

func fetchPath() string {
	var path string
	flag.StringVar(&path, "c", "local.yaml", "path to configs file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CFG_PATH")
	}

	return path
}
