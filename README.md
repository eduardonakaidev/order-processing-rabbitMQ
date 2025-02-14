# 🛒 Order Processing with RabbitMQ in Golang

This project demonstrates how to implement an **asynchronous order processing system** using **Golang** and **RabbitMQ**. The system consists of:

- **Publisher (API Server):** A REST API that receives orders and publishes messages to a RabbitMQ queue.
- **Worker (Consumer):** A worker that listens to the queue and processes orders asynchronously.
- **RabbitMQ:** Acts as the message broker to handle asynchronous communication.

## 🚀 Features
- Asynchronous processing of orders using RabbitMQ.
- Simple HTTP API using Golang's `net/http` package.
- Scalable architecture with multiple workers consuming messages.
- Supports multiple orders and error handling.

## 🏗️ Project Structure
```
order-processing-rabbitMQ/
│── publisher/                           # API Server (Publisher)
│   ├── cmd/main.go                     # Entry point for the publisher
│   ├── internal/server.go              # HTTP server setup
│   ├── internal/handler.go             # Handles HTTP requests
│   ├── models/order.go                 # Data models
│   ├── queue/rabbitmq.go               # RabbitMQ publisher logic
│   ├── queue/connection.go             # RabbitMQ publisher connection
│── worker/                              # Worker (Consumer)
│   ├── cmd/main.go                      # Entry point for the worker
│   ├── internal/queue/consumer.go       # Message consumer logic
│   ├── internal/queue/connection.go     # RabbitMQ connection manager
│── README.md                            # Project documentation
│── httpclient.http                      # HTTP request samples for testing
│── go.mod                               # Go module dependencies
│── go.sum                               # Go module checksum
```

## 🛠️ Setup & Installation

### 1️⃣ Start RabbitMQ (Docker)
Ensure RabbitMQ is running. If you don’t have it installed, you can run:

```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```
RabbitMQ management UI: [http://localhost:15672](http://localhost:15672) (user: `guest`, password: `guest`).

---

### 2️⃣ Install Dependencies
Go to the project directory and install the required dependencies:

```bash
go get github.com/streadway/amqp
```

---

### 3️⃣ Start the API Server (Publisher)
```bash
cd publisher
go run main.go
```
The API will start on **http://localhost:8080**.

---

### 4️⃣ Start the Worker (Consumer)
```bash
cd worker
go run main.go
```
The worker will listen for messages and process them asynchronously.

---

### 5️⃣ Test the API
Use `curl` or the `test_requests.http` file to send requests.

#### Send a new order:
```http
POST http://localhost:8080/order
Content-Type: application/json

{
  "id": "12345"
}
```

You should see the worker processing the order:
```
📦 Processing order: 12345
✅ Order 12345 processed successfully!
```

---

## 🛡️ Error Handling
- If RabbitMQ is down, the publisher/worker will log an error.
- Ensure RabbitMQ is running before starting the services.
- Check RabbitMQ logs with:
  ```bash
  docker logs rabbitmq
  ```

---

## 📌 Future Improvements
- ✅ Store processed orders in a database.
- ✅ Implement a **Dead Letter Queue (DLQ)** for failed messages.
- ✅ Add retries for message processing.
- ✅ Improve error handling and logging.

---

## 📜 License
This project is open-source and available under the **MIT License**.

---

🚀 **Happy Coding!** If you have any issues, feel free to ask. 🎯
