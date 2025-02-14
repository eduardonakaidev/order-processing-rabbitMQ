package main

import (
	"log"

	"github.com/eduardofrnkdev/order-processing-rabbitMQ/worker/internal/queue"
)

func main() {
	log.Println("🚀 Iniciando o worker...")

	// Criar conexão RabbitMQ
	rabbitMQ, err := queue.NewRabbitMQ()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Iniciar o consumidor
	err = queue.StartConsumer(rabbitMQ)
	if err != nil {
		log.Fatalf("❌ Erro ao iniciar consumidor: %v", err)
	}

	// Manter o processo rodando
	select {}
}
