package serv

import (
	"net/http"

	"github.com/cebilon123/gorl/internal/communication"
	"github.com/cebilon123/gorl/internal/config"
)

// CreateAndStartServer creates server instance and
// runs it.
func CreateAndStartServer(config config.Configer) error {
	semaphore := make(chan struct{}, config.MaxConcurrentRequests())
	webSocketModule := communication.NewWebsocketModule(semaphore)

	if err := webSocketModule.Start(); err != nil {
		return err
	}

	err := http.ListenAndServe(config.Port(), nil)

	if err != nil {
		return err
	}

	return nil
}
