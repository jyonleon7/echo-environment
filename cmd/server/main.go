package main

import (
	"log"

	"github.com/jyonleon7/echo-environment/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
