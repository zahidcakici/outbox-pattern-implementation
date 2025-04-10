
# Outbox Implementation

## Overview

This project implements the **Outbox Pattern**, a reliable way to ensure eventual consistency between a database and external systems (e.g., message queues, APIs). The pattern is commonly used in distributed systems to handle scenarios where operations span multiple services.

## Key Features

- **Transactional Integrity**: Ensures that database changes and message dispatches are atomic.
- **Retry Mechanism**: Automatically retries failed message deliveries.
- **Scalability**: Designed to handle high-throughput environments.
- **Extensibility**: Easily integrates with various message brokers or external systems.

## How It Works

1. **Database Transaction**: Changes are written to the database along with an "outbox" table entry containing the message to be sent.
2. **Message Dispatcher**: A background process reads the outbox table, sends messages to the external system, and marks them as processed.
3. **Error Handling**: Failed messages remain in the outbox table for retries.

## File Structure

``` plaintext
/outbox-implementation
├── src/
│   ├── invoice-service/
│   │   ├── kafka/
│   │   │   ├── consumer.js
│   │   ├── models/
│   │   │   ├── invoice.js
│   │   ├── index.js
│   │
│   ├─- order-service/
│   │   ├── handlers/
│   │   │   ├── order_handler.go
│   │   ├── models/
│   │   │   ├── order.go
│   │   ├── main.go
│   │   
│   ├── outbox-processor/
│   │   ├── handlers/
│   │   │   ├── outbox_handler.go
│   │   ├── models/
│   │   │   ├── order.go
│   │   ├── main.go
│
├── README.md
├── docker-compose.yml
```

## Technologies Used

- **Node.js**: For the invoice service and Kafka consumer.
- **Go**: For the order service and outbox processor.
- **Kafka**: For message brokering.
- **PostgreSQL**: For the order database.
- **MongoDB**: For the invoice database.
- **Docker**: For containerization and orchestration.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/zahidcakici/outbox-pattern-implementation.git
   cd outbox-implementation
   ```

2. Install dependencies:

   ```bash
   # For Node.js services
   cd invoice-service
   npm install

   # For Go services
   cd order-service
   go mod tidy

   cd outbox-processor
   go mod tidy
   ```

3. Set up the environment:

   ``` bash
   docker-compose up -d
   ```

4. Run the services:

   ```bash
   # For Node.js services
   cd invoice-service
   npm start
   # For Go services
   cd order-service
   go run main.go
   cd outbox-processor
   go run main.go
   ```

5. Test the implementation:
   - Use Postman or curl to send requests to the order service.
   - Check the outbox table in the database to see if messages are being processed.
   - Monitor the logs of the outbox processor to see the message dispatching process from http://localhost:8080.

## Contributing

Contributions are welcome! Please follow the guidelines in `CONTRIBUTING.md` (if available).

## License

This project is licensed under the MIT License. See `LICENSE` for details.
