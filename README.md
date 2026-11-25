# StartTech Application – Enterprise DevOps Assessment

## Project Overview

This repository contains the frontend and backend components of the StartTech application, prepared for an Enterprise CI/CD deployment on Amazon Elastic Kubernetes Service (EKS).

The application is a full-stack task management system built with React and Go. The project demonstrates containerisation, automated testing, continuous integration and preparation for cloud deployment.

## Application Architecture

The application uses the following technologies:

| Component | Technology |
|---|---|
| Frontend | React, TypeScript and Vite |
| Backend | Go (Golang) and Gin |
| Database | MongoDB Atlas |
| Caching | Amazon ElastiCache for Redis |
| Containerisation | Docker |
| Container Registry | Amazon Elastic Container Registry (ECR) |
| Orchestration | Amazon EKS |
| CI/CD | GitHub Actions |
| Infrastructure | Terraform |
| Content Delivery | Amazon CloudFront |
| Frontend Hosting | Amazon S3 |

## Repository Structure

```text
starttech-application/
│
├── .github/
│   └── workflows/
│       ├── backend-ci-cd.yml
│       └── frontend-ci-cd.yml
│
├── backend/
│   ├── cmd/
│   ├── internal/
│   ├── docs/
│   ├── Dockerfile
│   ├── .dockerignore
│   ├── .env.example
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── public/
│   ├── src/
│   ├── package.json
│   ├── package-lock.json
│   └── .env.example
│
├── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
│
├── scripts/
│   ├── deploy-backend.sh
│   ├── deploy-frontend.sh
│   ├── health-check.sh
│   └── rollback.sh
│
├── .gitignore
└── README.md
```

## CI/CD Workflows

GitHub Actions is used to automate application validation and build processes.

### Backend Pipeline

The backend workflow performs the following operations:

1. Checks out the repository.
2. Sets up the required Go version.
3. Downloads Go dependencies.
4. Runs Go tests.
5. Sets up Docker Buildx.
6. Builds the backend Docker image.

### Frontend Pipeline

The frontend workflow performs these operations:

1. Checks out the repository.
2. Sets up Node.js.
3. Installs dependencies using `npm ci`.
4. Runs ESLint.
5. Builds the React application.
6. Verifies the production build output.

The current workflows perform CI checks and image building. AWS deployment steps will be configured as the infrastructure and deployment components are completed.

## Local Development

### Backend

Requirements:
- Go
- MongoDB
- Redis (optional)

Navigate to the backend directory:

```bash
cd backend
```

Download dependencies:

```bash
go mod download
```

Create a local `.env` file using `.env.example` as a template, then configure the required database connection and authentication settings.

Run the application:

```bash
go run ./cmd/api
```

The backend listens on port `8080` by default.

### Frontend

Requirements:
- Node.js
- npm

Navigate to the frontend directory:

```bash
cd frontend
```

Install dependencies:

```bash
npm ci
```

Create a local `.env` file using `.env.example` as a template.

Start the development server:

```bash
npm run dev
```

Build the production application:

```bash
npm run build
```

Run code quality checks:

```bash
npm run lint
```

## Environment Configuration

Environment templates are provided in:

- `backend/.env.example`
- `frontend/.env.example`

Create local `.env` files from these templates and provide the required configuration values.

Never commit actual credentials, database connection strings, JWT secrets, AWS access keys or other sensitive information to GitHub.

Production secrets should be managed through appropriate secure services rather than committed to the repository.

## Kubernetes Deployment

The `k8s/` directory contains the Kubernetes deployment, service and ingress configuration files.

These manifests are intended for deployment to Amazon EKS as part of the assessment infrastructure.

## Infrastructure Repository

The Terraform infrastructure code is maintained separately in the `starttech-infra` repository.

It contains the AWS infrastructure configuration required for the application deployment.

## Project Status

This repository is being prepared for the StartTech Enterprise DevOps Assessment.

Current work includes:

- Application source code integration.
- Frontend API configuration.
- Backend API configuration.
- Dockerfile preparation.
- GitHub Actions CI workflow configuration.
- Kubernetes deployment manifest preparation.

AWS infrastructure deployment and production CI/CD integration are subsequent stages of the project.

## Security

- No production credentials should be committed.
- Environment files are excluded through `.gitignore`.
- Docker build context exclusions are configured in `.dockerignore`.
- Production secrets must be supplied securely during deployment.