package main

import (
	"log"
	"secgram/internal/server"
)

func main() {
	serv := server.NewServer()

	err := serv.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
