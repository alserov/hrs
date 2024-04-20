package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
	DB   Scylla `yaml:"db"`
}

type Scylla struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (s *Scylla) Dsn() string {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	if addr == "" {
		return fmt.Sprintf("%s:%s", os.Getenv("SCYLLA_MESSAGES_HOST"), os.Getenv("SCYLLA_MESSAGES_PORT"))
	}

	return addr
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
