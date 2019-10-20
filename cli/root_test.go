package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateConfig(t *testing.T) {
	t.Run("should return HttpConfig", func(t *testing.T) {
		expectedConfig := HTTPConfig{
			BaseURL: "https://api.github.com",
			URL:     "",
		}

		actualConfig := createConfig()
		assert.Equal(t, expectedConfig, actualConfig)
	})
}
