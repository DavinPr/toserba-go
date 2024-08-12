package log

import "github.com/rs/zerolog"

func Init(cfg Config) {
	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	zerolog.SetGlobalLevel(logLevel)
}
