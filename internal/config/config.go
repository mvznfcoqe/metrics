package config

import (
	"log"

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

	err := cleanenv.ReadConfig(".env", config)

	if err != nil {
		log.Fatal("Failed to load environment configuration")

		return nil, err
	}

	return config, nil
}
