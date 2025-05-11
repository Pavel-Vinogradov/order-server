package config

import (
	"github.com/spf13/viper"
	"log"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

type AppConfig struct {
	Name  string `yaml:"name"`
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

type Config struct {
	DBConfig *DBConfig `mapstructure:"database"`
	App      AppConfig `yaml:"app"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфига: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Ошибка разбора YAML: %v", err)
	}

	return &cfg
}
