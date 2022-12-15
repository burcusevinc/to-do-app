package main

import (
	"go.mod/server"
	"log"
)

func main() {
	// new server
	server := server.NewServer()
	// start server with port 3000
	err := server.StartServer(3000)
	if err != nil {
		log.Fatalln(err)
	}
}
