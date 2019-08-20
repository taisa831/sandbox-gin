package config

import "github.com/BurntSushi/toml"

type DBConfig struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
}

type Config struct {
	DB DBConfig `toml:"database"`
}

func NewConfig(path string, appMode string) (Config, error) {
	var conf Config
	confPath := path + appMode + ".toml"
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
