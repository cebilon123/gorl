package config_test

import (
	"os"
	"testing"

	"github.com/cebilon123/gorl/internal/config"
)

var (
	port                  = "1234"
	maxConcurrentRequests = "5"

	portKey                  = "PORT"
	maxConcurrentRequestsKey = "MAX_CONCURRENT_REQUESTS"
)

func TestEnvConfig_LoadConfig(t *testing.T) {
	type input struct {
		port                  string
		maxConcurrentRequests string
	}

	type expected struct {
		port                  string
		maxConcurrentRequests int
	}

	testCases := []struct {
		input    input
		expected expected
	}{
		{
			input: input{
				port:                  port,
				maxConcurrentRequests: maxConcurrentRequests,
			},
			expected: expected{
				port:                  ":" + port,
				maxConcurrentRequests: 5,
			},
		},
		{
			input: input{
				port:                  "",
				maxConcurrentRequests: "",
			},
			expected: expected{
				port:                  ":8080",
				maxConcurrentRequests: 15,
			},
		},
		{
			input: input{
				port:                  "asd",
				maxConcurrentRequests: "zxc",
			},
			expected: expected{
				port:                  ":8080",
				maxConcurrentRequests: 15,
			},
		},
	}

	for _, tt := range testCases {
		os.Setenv(portKey, tt.input.port)
		os.Setenv(maxConcurrentRequestsKey, tt.input.maxConcurrentRequests)

		envConfig := config.NewEnvConfig()
		if envConfig.MaxConcurrentRequests() != tt.expected.maxConcurrentRequests || envConfig.Port() != tt.expected.port {
			t.Errorf("NewEnvConfig: expected %v, actual: %v", tt.expected, envConfig)
		}
	}
}
