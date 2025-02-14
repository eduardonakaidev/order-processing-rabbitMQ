package server

import (
	"log"
	"net/http"

	"github.com/eduardofrnkdev/order-processing-rabbitMQ/publisher/config/queue"
)

var rabbitMQ *queue.RabbitMQConnection

// StartServer inicia o servidor HTTP e configura RabbitMQ
func StartServer() {
	var err error

	// Criar conexão RabbitMQ global
	rabbitMQ, err = queue.NewRabbitMQ()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Definir rotas
	http.HandleFunc("POST /order", orderHandler)

	log.Println("🚀 Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
