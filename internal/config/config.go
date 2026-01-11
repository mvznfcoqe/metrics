package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PrometheusURL    string   `env:"PROMETHEUS_URL" env-required:"true"`
	Port             string   `env:"PORT" env-required:"true"`
	AllowedOrigins   []string `env:"ALLOWED_ORIGINS" env-separator:"," env-default:"*"`
	CadvisorPort     string   `env:"CADVISOR_PORT" env-required:"true"`
	NodeExporterPort string   `env:"NODE_EXPORTER_PORT"`
}

func Load() (*Config, error) {
	config := &Config{}

	configPath := ".env"

	var loadConfigError error

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		loadConfigError = cleanenv.ReadEnv(config)
	} else {
		loadConfigError = cleanenv.ReadConfig(configPath, config)
	}

	if loadConfigError != nil {
		log.Fatal("Failed to load environment configuration")

		return nil, loadConfigError
	}

	return config, nil
}
