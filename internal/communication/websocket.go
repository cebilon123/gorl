package communication

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	start      = 0x1
	done       = 0x2
	canExecute = 0x3
)

// webSocket represents module used to
// do the communication based on
// web sockets
type webSocket struct {
	semaphore chan struct{}
	upgrader  websocket.Upgrader
}

func NewWebsocketModule(sem chan struct{}) Communicator {
	upg := websocket.Upgrader{}
	upg.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return webSocket{
		semaphore: sem,
		upgrader:  upg,
	}
}

func (ws webSocket) Start() error {
	log.Println("Starting web-socket communication module")
	http.HandleFunc(SemaphoreAddr, func(w http.ResponseWriter, r *http.Request) {
		c, err := ws.upgrader.Upgrade(w, r, nil)
		ctx := r.Context()
		ctx, cancel := context.WithCancel(ctx)

		log.Printf("Connected: %v", c.RemoteAddr())
		if err != nil {
			log.Print("Web-socket module | upgrade: ", err)
			cancel()
			return
		}

		defer c.Close()
		c.SetCloseHandler(func(code int, text string) error {
			log.Println("Web-socket module | connection closed by peer: ", text)
			cancel()
			return nil
		})

		for {
			select {
			case <-ctx.Done():
				return
			default:
				mt, _, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
				}
				if mt >= 1000 && mt <= 1015 {
					cancel()
				}

				err = c.WriteMessage(websocket.TextMessage, []byte("siema"))
				if err != nil {
					log.Println("write:", err)
				}
			}
		}
	})

	return nil
}
