package log

import "github.com/DavinPr/toserba-go/config"

type Config struct {
	LogLevel string
}

func NewConfig() Config {
	return Config{
		LogLevel: config.MustGetString("LOG_LEVEL"),
	}
}
