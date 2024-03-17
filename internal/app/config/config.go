package config

import "flag"

// Структура для конфигурации сервера
// Host - адрес сервера
// Prefix - базовый адрес результирующего сокращённого URL
type AppConfig struct {
	Host     string
	Location string
}

// Парсер конфигурации сервиса
func ParseFlags() *AppConfig {
	appConfig := AppConfig{Host: "", Location: ""}
	flag.Func("a", "Host:Port, default: localhost:8080", func(s string) error {
		if s != "" {
			appConfig.Host = s
		} else {
			appConfig.Host = "localhost:8080"
		}
		return nil
	})
	flag.Func("b", "Location, default: /", func(s string) error {
		if s != "" {
			appConfig.Location = "/" + s
		}
		return nil
	})

	flag.Parse()
	return &appConfig
}
