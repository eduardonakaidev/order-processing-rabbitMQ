package queue

import (
	"log"
	"time"


)

// StartConsumer inicia o consumidor do RabbitMQ
func StartConsumer(rabbitMQ *RabbitMQConnection) error {
	ch := rabbitMQ.GetChannel()

	// Declarar a fila (garante que a fila existe)
	q, err := ch.QueueDeclare(
		"orders_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Iniciar consumo de mensagens
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,  // Auto-Ack
		false, // Exclusivo
		false, // No-local
		false, // No-wait
		nil,
	)
	if err != nil {
		return err
	}

	// Goroutine para processar mensagens
	go func() {
		for d := range msgs {
			log.Printf("📦 Processando pedido: %s", d.Body)
			time.Sleep(2 * time.Second) // Simula um processamento
			log.Printf("✅ Pedido %s processado com sucesso!", d.Body)
		}
	}()

	log.Println("🎧 Worker aguardando mensagens...")
	return nil
}
