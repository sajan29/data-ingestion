# 🛠️ Data Ingestion Pipeline (Go + LocalStack)

A lightweight and modular Go-based data ingestion service that:
- Fetches data from a public API
- Transforms it by adding metadata (timestamp, source)
- Uploads the result as a JSON file to an AWS S3 bucket (emulated using LocalStack)

---

## 📦 Clone Repository

```bash
git clone https://github.com/sajan29/data-ingestion.git
cd data-ingestion
```

---

## 🔧 Setup

1. **Download dependencies**
```bash
go mod download
```

2. **Make localstack init script executable**
```bash
chmod +x localstack-init.sh
```

---

## 🚀 Running the Application

### ▶️ For Development (interpreted mode with hot reload)
```bash
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build
```

### 🏗️ For Production (compiled binary mode)
```bash
docker-compose -f docker-compose.yml up --build
```

---

## 🪣 S3 Bucket Checks (LocalStack)

### ✅ Check if bucket was created
```bash
docker exec -it localstack_s3_test awslocal s3 ls
```

### 📁 List objects inside the bucket
```bash
docker exec -it localstack_s3_test awslocal s3 ls s3://data-ingestion-bucket
```

### 📄 Download and view specific uploaded file
```bash
docker exec -it localstack_s3_test awslocal s3api get-object \
    --bucket data-ingestion-bucket \
    --key ingestion-20250611T165905.json \
    /tmp/ingestion-20250611T165905.json

docker exec -it localstack_s3_test cat /tmp/ingestion-20250611T165905.json
```

---

## 📝 Logs

```bash
docker logs data_ingestion_app
```

---

## ✅ Running Tests

```bash
go test ./test/...
```

---

## 📁 Project Structure

```
data-ingestion/
├── cmd/                   # Entry point for the app
├── internal/
│   ├── collector/         # API fetch logic
│   ├── transformer/       # Data transformation logic
│   ├── storage/           # AWS S3 upload logic
│   ├── models/            # Domain models
│   ├── config/            # Configuration management
│   └── utils/             # Logger and helper utilities
├── test/                  # All test cases
├── Dockerfile             # Docker build for prod
├── docker-compose.yml     # Base docker-compose config
├── docker-compose.dev.yml # Dev-specific overrides
├── localstack-init.sh     # LocalStack S3 bootstrap script
├── .env                   # Env variable definitions
├── go.mod / go.sum        # Go module definitions
└── README.md              # Project documentation
```

---

## 🧪 Tech Stack

- **Language:** Go
- **API Client:** `net/http`
- **AWS SDK:** v1
- **S3 (Mocked):** LocalStack
- **Containerization:** Docker + Docker Compose
- **Testing:** `go test`, `testify`