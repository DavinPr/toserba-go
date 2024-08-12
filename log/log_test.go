package log_test

import (
	"testing"

	"github.com/DavinPr/toserba-go/log"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func Test_InitWithValidLogLevel(t *testing.T) {
	cfg := log.Config{LogLevel: "info"}
	log.Init(cfg)
	assert.Equal(t, zerolog.InfoLevel, zerolog.GlobalLevel())
}

func Test_InitWithInvalidLogLevel(t *testing.T) {
	cfg := log.Config{LogLevel: "invalid"}
	assert.Panics(t, func() {
		log.Init(cfg)
	})
}

func Test_InitWithEmptyLogLevel(t *testing.T) {
	cfg := log.Config{LogLevel: "abc"}
	assert.Panics(t, func() {
		log.Init(cfg)
	})
}
