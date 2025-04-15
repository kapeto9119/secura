# Secura Project Structure

This document outlines the repository structure and organization for the Secura project.

## Repository Structure

```
secura/
├── api/                    # Go-based API Gateway
│   ├── cmd/
│   │   └── server/         # Application entry point
│   ├── internal/
│   │   ├── auth/           # Authentication and authorization
│   │   ├── blockchain/     # Blockchain integration logic
│   │   ├── config/         # Configuration handling
│   │   ├── handlers/       # API request handlers
│   │   ├── middlewares/    # HTTP middleware components
│   │   ├── models/         # Data models and structures
│   │   ├── services/       # Business logic services
│   │   └── storage/        # Data storage interfaces
│   ├── pkg/                # Reusable packages
│   │   └── anonymization/  # Client for anonymization service
│   ├── Dockerfile          # Container definition for API
│   ├── go.mod              # Go module definition
│   └── go.sum              # Go dependencies checksum
│
├── nlp/                    # Python NLP-based Anonymization Service
│   ├── app/
│   │   ├── api/            # FastAPI routes
│   │   ├── core/           # Application core
│   │   ├── models/         # NLP models and data structures
│   │   ├── services/       # Business logic services
│   │   └── utils/          # Utility functions
│   ├── tests/              # Unit and integration tests
│   ├── Dockerfile          # Container definition for NLP service
│   ├── pyproject.toml      # Python project configuration
│   └── requirements.txt    # Python dependencies
│
├── blockchain/             # Blockchain integration
│   ├── contracts/          # Smart contract definitions
│   │   ├── consent/        # Consent management contracts
│   │   └── audit/          # Audit trail contracts
│   ├── migrations/         # Contract deployment scripts
│   ├── test/               # Contract tests
│   └── truffle-config.js   # Truffle configuration
│
├── frontend/               # Next.js Admin Dashboard
│   ├── components/         # React components
│   ├── contexts/           # React contexts
│   ├── hooks/              # Custom React hooks
│   ├── pages/              # Next.js pages
│   ├── public/             # Static assets
│   ├── styles/             # CSS/SCSS styles
│   ├── utils/              # Utility functions
│   ├── Dockerfile          # Container definition for frontend
│   ├── next.config.js      # Next.js configuration
│   └── package.json        # Node.js dependencies
│
├── db/                     # Database migrations and schemas
│   ├── migrations/         # SQL migration files
│   └── init.sql            # Initial database setup
│
├── docs/                   # Documentation
│   ├── architecture/       # Architecture diagrams and docs
│   ├── api/                # API documentation
│   └── guides/             # User and developer guides
│
├── scripts/                # Development and deployment scripts
│   ├── dev/                # Local development helpers
│   └── deploy/             # Deployment automation
│
├── .github/                # GitHub configuration
│   └── workflows/          # GitHub Actions CI/CD workflows
│
├── docker-compose.yml      # Docker Compose for local development
├── .gitignore              # Git ignore patterns
├── README.md               # Project overview and documentation
├── DEVELOPMENT_PLAN.md     # Development roadmap and progress
└── LICENSE                 # Project license
```

## Component Details

### API Gateway (Go)

The central component that handles client requests, coordinates with other services, and manages authentication.

- **Key Responsibilities**:
  - Authentication and authorization
  - Request validation
  - Routing to appropriate services
  - LLM provider integration
  - Response handling

- **Technologies**:
  - Go (1.20+)
  - Gin Web Framework
  - JWT for authentication
  - gRPC for internal service communication

### Anonymization Service (Python)

NLP-based service for detecting and masking personally identifiable information (PII) in text.

- **Key Responsibilities**:
  - Identify sensitive information in text
  - Replace or mask PII with appropriate tokens
  - Ensure context preservation
  - Handle domain-specific anonymization (e.g., healthcare)

- **Technologies**:
  - Python (3.9+)
  - FastAPI
  - spaCy or Hugging Face Transformers for NER
  - Docker for containerization

### Blockchain Integration

Manages the immutable audit trail using blockchain technology.

- **Key Responsibilities**:
  - Create and manage blockchain transactions
  - Generate and verify cryptographic hashes
  - Provide audit trail verification
  - Handle consent management (future)

- **Technologies**:
  - Ethereum-compatible blockchain (for MVP: private network)
  - Solidity for smart contracts
  - Web3.js/ethers.js for interaction
  - Truffle for development and testing

### Admin Dashboard (Next.js)

Web interface for monitoring, compliance verification, and system management.

- **Key Responsibilities**:
  - Display audit logs
  - Provide search and filtering capabilities
  - Show anonymization details
  - Admin user management

- **Technologies**:
  - Next.js (React framework)
  - TypeScript
  - Tailwind CSS for styling
  - React Query for data fetching

### Database

Stores metadata, configuration, and references to blockchain records.

- **Key Responsibilities**:
  - Store user accounts and permissions
  - Maintain configuration data
  - Store references to blockchain transactions
  - Cache query results for performance

- **Technologies**:
  - PostgreSQL 13+
  - Migrations using golang-migrate or similar

## Implementation Priority

For the MVP, we'll focus on implementing these components in the following order:

1. API Gateway - Basic structure and OpenAI integration
2. Anonymization Service - Simple NER-based PII detection
3. Blockchain Integration - Basic audit trail functionality
4. Admin Dashboard - Minimal viable interface

## Getting Started

Instructions for setting up the development environment will be added as each component is implemented. 