package config

type Config struct {
	DatabaseConnection string `yaml:"db"`
	AppPort            string `yaml:"port"`
}
