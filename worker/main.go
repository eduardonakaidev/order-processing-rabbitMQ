package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

// Conectar ao RabbitMQ e consumir mensagens
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Erro ao conectar ao RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erro ao abrir um canal: %s", err)
	}
	defer ch.Close()

	// Declarar a fila
	q, err := ch.QueueDeclare(
		"orders_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Erro ao declarar fila: %s", err)
	}

	// Consumir mensagens
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Erro ao registrar consumidor: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("📦 Processando pedido: %s", d.Body)
			time.Sleep(2 * time.Second) // Simula um processamento
			log.Printf("✅ Pedido %s processado com sucesso!", d.Body)
		}
	}()

	log.Println("🎧 Worker aguardando mensagens...")
	<-forever
}
