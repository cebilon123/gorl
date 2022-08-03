package communication

// Communicator must be implemented by all
// structs which are used for communication
// between client and server
type Communicator interface {
	// Start communication module. Error
	// will be nil, if there is none
	Start() error
}

const (
	// SemaphoreAddr is a endpoint
	// addres to start communication
	// between client and server
	SemaphoreAddr = "/semaphore"
)
