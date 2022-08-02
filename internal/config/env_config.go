package config

import (
	"os"
	"strconv"
)

type envConfig struct {
	port                  string
	maxConcurrentRequests int
}

const (
	portKey                  = "PORT"
	maxConcurrentRequestsKey = "MAX_CONCURRENT_REQUESTS"

	defaultPort                  = "8080"
	defaultConcurrentRequests    = "15"
	defaultConcurrentRequestsInt = 15
)

func NewEnvConfig() Configer {
	port, exists := os.LookupEnv(portKey)
	_, err := strconv.Atoi(port)
	if !exists || err != nil {
		port = defaultPort
	}
	port = ":" + port

	maxConcurrentRequests, exists := os.LookupEnv(maxConcurrentRequestsKey)
	_, err = strconv.Atoi(maxConcurrentRequests)
	if !exists || err != nil {
		maxConcurrentRequests = defaultConcurrentRequests
	}

	parsedConcurrentRequests, err := strconv.Atoi(maxConcurrentRequests)
	if err != nil {
		parsedConcurrentRequests = defaultConcurrentRequestsInt
	}

	return envConfig{
		port:                  port,
		maxConcurrentRequests: parsedConcurrentRequests,
	}
}

func (c envConfig) Port() string {
	return c.port
}

func (c envConfig) MaxConcurrentRequests() int {
	return c.maxConcurrentRequests
}
