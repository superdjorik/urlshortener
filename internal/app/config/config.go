package config

import "flag"

// Структура для конфигурации сервера
// Host - адрес сервера
// Prefix - базовый адрес результирующего сокращённого URL
var options struct {
	host     string
	location string
}

// Парсер конфига
func ParseFlags() {
	flag.StringVar(&options.host, "a", ":8080", "address and port to run server")
	flag.StringVar(&options.location, "b", "http://localhost:8080/", "server address")
	flag.Parse()
}

func Host() string {
	return options.host
}

func Location() string {
	return options.location
}
