package queue

import (
	"fmt"

	"sync"

	"github.com/streadway/amqp"
)

// RabbitMQConnection estrutura para gerenciar a conexão com RabbitMQ
type RabbitMQConnection struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	mu   sync.Mutex
}

// NewRabbitMQ cria uma nova instância da conexão RabbitMQ
func NewRabbitMQ() (*RabbitMQConnection, error) {
	// Conectar ao RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao RabbitMQ: %w", err)
	}

	// Criar canal
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("erro ao abrir um canal: %w", err)
	}

	return &RabbitMQConnection{
		conn: conn,
		ch:   ch,
	}, nil
}

// GetChannel retorna o canal ativo do RabbitMQ
func (r *RabbitMQConnection) GetChannel() *amqp.Channel {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.ch
}

// Close fecha a conexão e o canal do RabbitMQ
func (r *RabbitMQConnection) Close() {
	r.ch.Close()
	r.conn.Close()
}
