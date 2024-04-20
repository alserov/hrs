package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
	DB   DB     `yaml:"db"`
}

type DB struct {
	Scylla `yaml:"scylla"`
}

type Scylla struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Keyspace string `yaml:"keyspace"`
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
		panic("failed to parse config: " + err.Error())
	}

	return &cfg
}

func fetchPath() string {
	var path string
	flag.StringVar(&path, "c", "local.yaml", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CFG_PATH")
	}

	return path
}
