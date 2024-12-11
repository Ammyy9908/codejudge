# Backend System for Online Code Compiler

This project implements an online code compiler system with a microservices architecture. The system includes the following key services:

1. **Submission Service**: Handles user code submissions and pushes them to a RabbitMQ queue.
2. **Execution Service**: Executes the submitted code in a sandboxed environment and publishes the results.
3. **Notification Service**: Sends execution results back to users via WebSocket.
4. **WebSocket Service**: Manages WebSocket connections and relays notifications to clients.

---

## **Folder Structure**

```plaintext
backend/
├── cmd/                        # Entry points for services
│   ├── submission-service/     # Submission service entry point
│   │   └── main.go
│   ├── execution-service/      # Execution service entry point
│   │   └── main.go
│   ├── notification-service/   # Notification service entry point
│   │   └── main.go
│   └── websocket-service/      # WebSocket service entry point
│       └── main.go
├── internal/                   # Service-specific logic
│   ├── submission/             # Logic for submission-service
│   │   ├── handlers/           # HTTP handlers
│   │   ├── queue/              # RabbitMQ producer for submissions
│   │   └── models.go           # Submission models
│   ├── execution/              # Logic for execution-service
│   │   ├── sandbox/            # Code execution sandbox
│   │   ├── queue/              # RabbitMQ consumer for submissions
│   │   └── models.go           # Execution result models
│   ├── notification/           # Logic for notification-service
│   │   ├── websocket/          # WebSocket integration
│   │   ├── queue/              # RabbitMQ consumer for execution results
│   │   └── models.go           # Notification models
│   └── websocket/              # Logic for WebSocket service
│       ├── handlers/           # WebSocket connection handlers
│       ├── redis/              # Redis utilities
│       └── models.go           # WebSocket models
├── pkg/                        # Shared libraries and utilities
│   ├── rabbitmq/               # RabbitMQ connection and utility functions
│   ├── redis/                  # Redis connection and utility functions
│   ├── logger/                 # Centralized logging utilities
│   ├── config/                 # Configuration utilities
│   └── utils/                  # Miscellaneous utilities (e.g., UUID generator)
├── deployments/                # Deployment configurations
│   ├── k8s/                    # Kubernetes YAML manifests
│   │   ├── submission-deployment.yaml
│   │   ├── execution-deployment.yaml
│   │   ├── notification-deployment.yaml
│   │   └── websocket-deployment.yaml
│   └── docker-compose.yaml     # Local setup for Docker Compose
├── tests/                      # Integration and unit tests
│   ├── submission/             # Tests for submission service
│   ├── execution/              # Tests for execution service
│   ├── notification/           # Tests for notification service
│   └── websocket/              # Tests for WebSocket service
└── README.md                   # Documentation for the backend system
```

---

## **System Workflow**

### **1. Submission**
- The frontend sends code and metadata (language, user ID) to the Submission Service via an API.
- The Submission Service assigns a unique correlation ID to the task and pushes it to a RabbitMQ `submission_queue`.

### **2. Execution**
- The Execution Service consumes tasks from `submission_queue`.
- It executes the code in a sandboxed environment (e.g., Docker) and generates results.
- Results are published to a RabbitMQ `execution_queue`.

### **3. Notification**
- The Notification Service listens to `execution_queue` for results.
- Results are sent to the appropriate WebSocket connection via the WebSocket Service.

### **4. WebSocket Communication**
- The WebSocket Service manages user connections and ensures results are relayed to the correct client.

---

## **Setup and Installation**

### **Prerequisites**
- **Docker** and **Docker Compose** for local development.
- **RabbitMQ** for message queuing.
- **Redis** for managing WebSocket state.
- **Golang** for backend services.

### **Steps**

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd backend
   ```

2. Start the services using Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Access RabbitMQ management interface:
   - URL: `http://localhost:15672`
   - Username: `guest`, Password: `guest`

4. Start each service locally:
   ```bash
   go run cmd/submission-service/main.go
   go run cmd/execution-service/main.go
   go run cmd/notification-service/main.go
   go run cmd/websocket-service/main.go
   ```

---

## **Configuration**

- All service configurations (e.g., RabbitMQ, Redis) are stored in `config/`.
- Example RabbitMQ configuration in `config/rabbitmq.yaml`:
  ```yaml
  rabbitmq:
    url: amqp://guest:guest@localhost:5672/
    submission_queue: submission_queue
    execution_queue: execution_queue
  ```

---

## **Testing**

- Run unit tests:
  ```bash
  go test ./...
  ```

- Run integration tests:
  ```bash
  cd tests
  go test -v
  ```

---

## **Deployment**

### **Docker Compose**
For local development, use the `docker-compose.yaml` file:

```bash
docker-compose up --build
```

### **Kubernetes**
For production, deploy the services using Kubernetes manifests in `deployments/k8s/`:

```bash
kubectl apply -f deployments/k8s/
```

---

## **Future Enhancements**

1. Implement **load balancing** for WebSocket connections.
2. Add **authentication and authorization** for APIs and WebSocket connections.
3. Introduce **code execution quotas** to prevent abuse.
4. Add **metrics and monitoring** using Prometheus and Grafana.
5. Improve sandboxing using Firecracker or gVisor for enhanced security.

---

## **License**

This project is licensed under the MIT License. See the `LICENSE` file for details.
