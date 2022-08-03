package main

import (
	"log"

	"github.com/cebilon123/gorl/internal/config"
	"github.com/cebilon123/gorl/internal/serv"
)

func main() {
	config := config.NewEnvConfig()
	err := serv.CreateAndStartServer(config)

	if err != nil {
		log.Panic(err)
	}
}
