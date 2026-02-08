package config

import "os"

// Config armazena as configurações da aplicação
type Config struct {
	Port        string
	Environment string
}

// Load carrega as configurações da aplicação
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

// getEnv obtém o valor de uma variável de ambiente com um valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
