package queue

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

// PublishMessage publica uma mensagem no RabbitMQ usando uma conexão existente
func PublishMessage(rabbitMQ *RabbitMQConnection, orderID string) error {
	ch := rabbitMQ.GetChannel()

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
		return fmt.Errorf("erro ao declarar fila: %w", err)
	}

	// Publicar a mensagem
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(orderID),
		},
	)
	if err != nil {
		return fmt.Errorf("erro ao publicar mensagem: %w", err)
	}

	log.Printf("📦 Pedido enviado para a fila: %s", orderID)
	return nil
}
