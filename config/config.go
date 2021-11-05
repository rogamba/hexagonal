package config

import "os"

type AppConfig struct {
	Env           string
	Port          string
	RedisEndpoint string
}

func ParseAppConfig() (AppConfig, error) {
	return AppConfig{
		Env:  getenv("ENV", "dev"),
		Port: getenv("PORT", "1234"),
	}, nil
}

func getenv(key, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}
	return val
}
