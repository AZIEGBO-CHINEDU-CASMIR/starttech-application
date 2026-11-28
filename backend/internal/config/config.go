package config

import (
    "strings"

    "github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
    ServerPort         string   `mapstructure:"PORT"`
    MongoURI           string   `mapstructure:"MONGO_URI"`
    DBName             string   `mapstructure:"DB_NAME"`
    JWTSecretKey       string   `mapstructure:"JWT_SECRET_KEY"`
    JWTExpirationHours int      `mapstructure:"JWT_EXPIRATION_HOURS"`
    EnableCache        bool     `mapstructure:"ENABLE_CACHE"`
    RedisAddr          string   `mapstructure:"REDIS_ADDR"`
    RedisPassword      string   `mapstructure:"REDIS_PASSWORD"`
    RedisTLS            bool     `mapstructure:"REDIS_TLS"`
    LogLevel            string   `mapstructure:"LOG_LEVEL"`
    LogFormat           string   `mapstructure:"LOG_FORMAT"`
    CookieDomains       []string `mapstructure:"COOKIE_DOMAINS"`
    SecureCookie        bool     `mapstructure:"SECURE_COOKIE"`
    AllowedOrigins      []string `mapstructure:"ALLOWED_ORIGINS"`
}

// cleanList splits comma-separated values and removes surrounding spaces and quotes.
func cleanList(values []string) []string {
    var cleaned []string

    for _, value := range values {
        for _, part := range strings.Split(value, ",") {
            trimmed := strings.TrimSpace(part)
            trimmed = strings.Trim(trimmed, "\"'")
            if trimmed != "" {
                cleaned = append(cleaned, trimmed)
            }
        }
    }

    return cleaned
}

// LoadConfig reads configuration from a .env file or environment variables.
func LoadConfig(path string) (config Config, err error) {
    v := viper.New()

    v.AddConfigPath(path)
    v.SetConfigName(".env")
    v.SetConfigType("env")
    v.AutomaticEnv()

    // Explicitly bind every environment variable.
    keys := []string{
        "PORT",
        "MONGO_URI",
        "DB_NAME",
        "JWT_SECRET_KEY",
        "JWT_EXPIRATION_HOURS",
        "ENABLE_CACHE",
        "REDIS_ADDR",
        "REDIS_PASSWORD",
        "REDIS_TLS",
        "LOG_LEVEL",
        "LOG_FORMAT",
        "COOKIE_DOMAINS",
        "SECURE_COOKIE",
        "ALLOWED_ORIGINS",
    }

    for _, key := range keys {
        if err = v.BindEnv(key); err != nil {
            return
        }
    }

    // Default values.
    v.SetDefault("PORT", "8080")
    v.SetDefault("ENABLE_CACHE", false)
    v.SetDefault("REDIS_TLS", false)
    v.SetDefault("JWT_EXPIRATION_HOURS", 72)
    v.SetDefault("COOKIE_DOMAINS", []string{"localhost"})
    v.SetDefault("SECURE_COOKIE", false)
    v.SetDefault("ALLOWED_ORIGINS", []string{"http://localhost:5173"})
    v.SetDefault("LOG_LEVEL", "INFO")
    v.SetDefault("LOG_FORMAT", "json")

    // Read .env if available. Environment variables take precedence.
    err = v.ReadInConfig()
    if err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return
        }
    }

    err = v.Unmarshal(&config)
    if err != nil {
        return
    }

    config.AllowedOrigins = cleanList(v.GetStringSlice("ALLOWED_ORIGINS"))
    config.CookieDomains = cleanList(v.GetStringSlice("COOKIE_DOMAINS"))

    return
}
