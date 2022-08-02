package config

// Configer is an interface used to
// handle configs. Each configuration structure should
// implement it.
type Configer interface {
	Port() string
	MaxConcurrentRequests() int
}
