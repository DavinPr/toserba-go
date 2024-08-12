package log_test

import (
	"os"
	"testing"

	"github.com/DavinPr/toserba-go/log"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	configVars := map[string]string{
		"LOG_LEVEL": "info",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	cfg := log.NewConfig()
	assert.Equal(t, "info", cfg.LogLevel)
}
