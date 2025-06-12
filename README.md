# loan-buddy

> A Go microservice that simulates loan offer aggregation from multiple providers. Since this is a personal project the data is mocked.

## Features

* GET /offers – returns aggregated, sorted mock loan offers

* POST /offers – accepts a new loan offer (demo purposes)

* Exposes API via Echo framework

* Dockerized & deployed to Kubernetes (Kind)

* Helm chart for cloud-native packaging

* GitHub Actions CI pipeline

* Simulated GCP infrastructure with Terraform

---

## API Endpoints

### GET /offers

Returns aggregated offers from 3 mock providers, sorted by interest rate.

```json
[
  { "provider": "ProviderB", "rate": 2.9, "amount": 12000 },
  { "provider": "ProviderA", "rate": 3.5, "amount": 10000 },
  ...
]
```

### POST /offers

Accepts a JSON offer and echoes it back.

```shell
curl -X POST http://localhost:8080/offers \
  -H "Content-Type: application/json" \
  -d '{"provider":"TestProvider","rate":3.7,"amount":8000}'
```

---

## Running Locally

### Run with Go

```shell
# Run application
go run cmd/loanbuddy/main.go
# Run tests
go test ./...
# Run tests with coverage
go test ./... -cover
```

### Run with Docker

```shell
docker build -t loan-buddy .
docker run -p 8080:8080 loan-buddy
```

### Run in Kubernetes (Kind)

```shell
# Build and load into kind
docker build -t loan-buddy .
kind load docker-image loan-buddy

# Install with Helm
helm install loan-buddy ./helm/loan-buddy

# Access via port-forward
kubectl port-forward svc/loan-buddy 8080:8080
curl http://localhost:8080/offers
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

* [ ] Replace mock data with actual third-party APIs
* [ ] Store data in a real database
* [ ] Add authentication
* [ ] Do an actual deployemnt to GKE

---

## Feedback

This project is personal, but contributions, suggestions and feedback are welcome!
