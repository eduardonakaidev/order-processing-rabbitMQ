package main

import (
	"log"
	"github.com/eduardofrnkdev/order-processing-rabbitMQ/publisher/internal/server"
)

func main() {
	log.Println("🚀 Iniciando o servidor...")
	server.StartServer()
}
