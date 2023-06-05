package config

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/rs/zerolog/log"

	"gopkg.in/yaml.v2"
)

var Conf *Configs

// Configs : struct Configs
type Configs struct {
	Srv   Server `yaml:"srv"`
	Mysql Mysql  `yaml:"mysql"`
	Logs  Logs   `yaml:"logs"`
}

// Server : struct Server
type Server struct {
	Addr        string `yaml:"addr"`
	Certificate string `yaml:"certificate"`
	Key         string `yaml:"key"`
}

// Logs : struct Logs
type Logs struct {
	Level string `yaml:"level"`
}

// Mysql : struct Mysql
type Mysql struct {
	DSN string `yaml:"dsn"`
}

// New : func new
func New(file string) *Configs {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal().Err(err).Msgf("Homedir error")
	}

	path := filepath.Join(home, ".config", "api-go", file)

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal().Err(err).Msgf("Read config error")
	}

	var cfg Configs
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatal().Err(err).Msgf("Unmarshal config error")
	}

	return &cfg
}

// CheckRealEstate check if mandatory config is ok
func (c *Configs) CheckConfig() {
	if reflect.DeepEqual(c.Srv, Server{}) {
		log.Fatal().Msgf("Missing server config")
	}

	if reflect.DeepEqual(c.Mysql, Mysql{}) {
		log.Fatal().Msgf("Missing mysql config")
	}
}
