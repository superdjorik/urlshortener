package config

import "flag"

// Структура для конфигурации сервера
// Host - адрес сервера
// Prefix - базовый адрес результирующего сокращённого URL
type AppConfig struct {
	Host   string
	Prefix string
}

// Парсер конфигурации сервиса
func ParseFlags() *AppConfig {
	appConfig := AppConfig{Prefix: ""}
	flag.StringVar(&appConfig.Host, "a", ":8080", "Default Host:port")
	flag.Func("b", "App prefix", func(s string) error {
		if s != "" {
			appConfig.Prefix = "/" + s
		}
		return nil
	})

	flag.Parse()
	return &appConfig
}
