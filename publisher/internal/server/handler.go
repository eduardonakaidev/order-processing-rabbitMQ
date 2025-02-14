package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eduardofrnkdev/order-processing-rabbitMQ/publisher/config/queue"
	"github.com/eduardofrnkdev/order-processing-rabbitMQ/publisher/internal/models"
)

// orderHandler processa pedidos recebidos via HTTP
func orderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Enviar pedido para RabbitMQ usando a conexão existente
	err = queue.PublishMessage(rabbitMQ, order.ID)
	if err != nil {
		http.Error(w, "Erro ao processar pedido", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Pedido recebido", "order_id": "%s"}`, order.ID)
}
