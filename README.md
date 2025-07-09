# loan-buddy

> A Go microservice that simulates loan offer aggregation from multiple providers. Since this is a personal project the data is mocked.

## Features

* Go service:
  * GET `/offers` – returns aggregated, sorted mock loan offers
  * POST `/checkout` → initiates order and calls Python  `confirm-payment` API
* Python FastAPI:
  * POST `/confirm-payment` → returns `APPROVED` / `REJECTED`
* Dockerized & deployed to Kubernetes (Kind)
* Helm chart for cloud-native packaging
* GitHub Actions CI pipeline
* Simulated GCP infrastructure with Terraform

---

## Architecture

The system consists of two main services:

* **Go Backend**: Handles loan offers, user checkout.
* **Python FastAPI Service**: Simulates external confirmation (eg. payment)

### User Flow

In the current design, the Go service directly calls the Python FastAPI service over HTTP to receive approval status in real time.

![Current Architecture](./docs/architecture_current.png)

This synchronous design is simple and effective for interactive use cases, where the user expects an immediate response (eg. checkout confirmation).

However, in a scalable production system, this flow would be refactored into a decoupled architecture using asynchronous communication:

* The Go service publishes a `checkout_submitted` event to a queue like Google Pub/Sub.
* The Python service subscribes to this event stream, performs async scoring and writes status to a data store.
* User can:
  * Poll status
  * Or receive a notification (eg. email)

![Future Architecture](./docs/architecture_future.png)

---

## API Endpoints

### GO Backend: GET /offers

Returns aggregated offers from 3 mock providers, sorted by interest rate.

```shell
curl -X GET http://localhost:8080/offers 
```

```json
[
  { "id": "offer2", "provider": "ProviderB", "rate": 2.9, "amount": 12000 },
  { "id": "offer2", "provider": "ProviderA", "rate": 3.5, "amount": 10000 },
  ...
]
```

### GO Backend: POST /checkout

Submits a selected loan offer for processing. Sends a confirmation request to the Python service and returns the approval status.

```shell
curl -X POST http://localhost:8080/checkout \
  -H "Content-Type: application/json" \
  -d '{
        "user_id": "cristina",
        "loan_id": "offer2"
      }'
```

```json
{
  "id": "order-123",
  "user_id": "cristina",
  "loan_id": "offer2",
  "status": "APPROVED",
  "created_at": "2025-07-08T12:30:45Z"
}
```

### Python FastAPI: POST /confirm-payment

This endpoint is called internally by the Go backend after checkout. It simulates an external payment/approval system and returns a decision.

```shell
curl -X POST http://localhost:8000/confirm-payment \
  -H "Content-Type: application/json" \
  -d '{
        "order_id": "order-123",
        "amount": 12000,
        "status": "PENDING"
      }'
```

```json
{
  "status": "APPROVED"
}
```

---

## Running Locally

### Run Go Service

```shell
cd go/
# Run application
go run cmd/loanbuddy/main.go
# Run tests
go test ./...
# Run tests with coverage
go test ./... -cover
```

### Run Python Service

```shell
cd python
# Install dependencies
pip install -r requirements.txt
# Start the FastAPI server
uvicorn main:app --reload --port 800
```

### Run with Docker

Go Service:

```shell
cd go/
docker build -t loan-buddy .
docker run -p 8080:8080 loan-buddy
```

Python Service:

```shell
cd python/
docker build -t loan-buddy-python -f python/Dockerfile ./python
docker run -p 8000:8000 loan-buddy-python
```

---

## CI/CD

### Github Actions

The Github Actions workflow runs all tests and builds the Docker image. Check out `.github/workflow/ci.yaml`.

### Deployment to GCP using Terraform

The `terraform/` folder defines how a GKE cluster would be provisioned in GCP, simulationg the infrastructure that would host the Loan Buddy microservice.

```shell
cd terraform
terraform init
terraform plan
```

Note: No credentials are needed because the actual GKE cluster is never deployed in order to keep the project free.

---

## Ideas

To improve project:

* [x] Add Python FastAPI for checkout confirmation  
* [ ] Replace mock data with actual third-party APIs
* [ ] Store data in a persistent database
* [ ] Move approval to async (Pub/Sub) for decouplin
* [ ] Do an actual deployemnt to GKE
* [ ] Add retry queue for failed approval requests
* [ ] Export events to BigQuery for product analytics
* [ ] Add authentication

---

## Feedback

This project is personal, but contributions, suggestions and feedback are welcome!
