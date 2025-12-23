package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort  string `mapstructure:"APP_PORT"`
	LogLevel string `mapstructure:"LOG_LEVEL"`

	// Database Config
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	// SMTP Config
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPEmail    string `mapstructure:"SMTP_EMAIL"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
}

func LoadConfig() (*Config, error) {
	// Debug: Print CWD to help diagnose where the app is running from
	if wd, err := os.Getwd(); err == nil {
		log.Printf("Current Working Directory: %s", wd)
	}

	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "info")

	// DB Defaults
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_NAME", "postgres")

	// SMTP Defaults (Gmail)
	viper.SetDefault("SMTP_HOST", "smtp.gmail.com")
	viper.SetDefault("SMTP_PORT", 587)

	// Try to load .env from current directory
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not load .env from current directory: %v", err)

		// Fallback: Try to load from parent directory (common when debugging from cmd/api)
		viper.SetConfigFile("../.env")
		if err := viper.ReadInConfig(); err != nil {
			// Fallback: Try to load from cmd/api directory (if running from root but file is in cmd/api)
			viper.SetConfigFile("cmd/api/.env")
			if err := viper.ReadInConfig(); err != nil {
				log.Printf("Warning: Could not load .env from ., ../, or cmd/api/.env: %v", err)
				log.Println("Using environment variables or defaults.")
			} else {
				log.Println("Loaded .env from cmd/api/.env")
			}
		} else {
			log.Println("Loaded .env from parent directory")
		}
	} else {
		log.Println("Loaded .env from current directory")
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
