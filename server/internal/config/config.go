package config

type Config struct {
	DatabaseConnection string `yaml:"db"`
	TokenKey           string `yaml:"token_key"`
	AppPort            string `yaml:"port"`
}
