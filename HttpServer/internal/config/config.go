// Package config centralizes environment-variable-based configuration —
// the Go equivalent of reading process.env in Node, just explicit and
// typed instead of implicit and stringly-typed everywhere.
package config

import "os"

type Config struct {
	Port   string
	APIKey string
}

// Load reads config from environment variables, falling back to sane
// defaults for local development. In production you'd set PORT/API_KEY
// via your deployment platform instead of relying on the defaults.
func Load() Config {
	return Config{
		Port:   getEnv("PORT", "8080"),
		APIKey: getEnv("API_KEY", "dev-secret-key"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
