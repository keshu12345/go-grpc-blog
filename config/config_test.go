package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	config := GetConfig("./local")
	assert.Equal(t, "local", config.EnvironmentName)
	assert.Equal(t, "localhost", config.Server.Host)
	assert.Equal(t, 50051, config.Server.Port)
}
