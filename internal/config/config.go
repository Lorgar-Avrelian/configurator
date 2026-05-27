package config

import (
	log "filler/internal/logger"
	"os"

	"gopkg.in/yaml.v3"
)

var cfg *Config

type Config struct {
	App    AppConfig    `yaml:"app"`
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"database"`
	Logger LoggerConfig `yaml:"logger"`
}

type AppConfig struct {
	Name string `yaml:"name"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DBConfig struct {
	URL      string `yaml:"url"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"name"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

func Init(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Не удалось открыть файл %s: %v", path, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Ошибка закрытия файла %s: %v", path, err)
		}
	}(file)

	var localCfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&localCfg); err != nil {
		log.Fatalf("Ошибка чтения файла %s: %v", path, err)
	}
	cfg = &localCfg
}

func Get() *Config {
	return cfg
}
