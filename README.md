# Loyalty Card Service

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.18+-blue)](https://golang.org/dl/)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/patricksferraz/loyalty-card/actions)
[![Docker](https://img.shields.io/badge/docker-ready-blue)](https://www.docker.com/)

---

## Table of Contents
- [Loyalty Card Service](#loyalty-card-service)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Demo \& API Docs](#demo--api-docs)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Quickstart (Docker Compose)](#quickstart-docker-compose)
    - [Manual Setup (Go)](#manual-setup-go)
    - [Environment Variables](#environment-variables)
  - [Usage](#usage)
    - [API Endpoints](#api-endpoints)
      - [Guests](#guests)
      - [Scores](#scores)
      - [Tags](#tags)
    - [Example: Create a Guest](#example-create-a-guest)
    - [API Documentation](#api-documentation)
  - [Development](#development)
    - [Makefile Commands](#makefile-commands)
    - [Running Tests](#running-tests)
    - [Contributing](#contributing)
  - [Tech Stack](#tech-stack)
  - [License](#license)
  - [Contact \& Community](#contact--community)
  - [Acknowledgements](#acknowledgements)

---

## Overview

**Loyalty Card** is a modern, open-source RESTful service for managing loyalty programs, guests, scores, and tags. Built with Go, Fiber, and GORM, it offers a robust API, modular architecture, and full Docker support. Ideal for businesses, developers, and open-source contributors looking for a scalable and extensible loyalty management backend.

---

## Demo & API Docs

- **Swagger/OpenAPI Documentation:** [View API Docs](app/rest/docs/swagger.yaml)
- **Live Swagger UI:** `/api/v1/swagger/index.html` (when running the service)

---

## Features
- **RESTful API** for managing guests, scores, and tags
- **Modular Go codebase** with clear separation of concerns
- **CLI support** via Cobra
- **OpenAPI/Swagger** auto-generated documentation
- **Dockerized** for easy deployment
- **Makefile** for streamlined development workflows
- **Environment-based configuration**
- **PostgreSQL & SQLite** support via GORM
- **Extensible and open for contributions**

---

## Getting Started

### Prerequisites
- [Go 1.18+](https://golang.org/dl/)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)

### Quickstart (Docker Compose)
```sh
git clone https://github.com/patricksferraz/loyalty-card.git
cd loyalty-card
cp .env.example .env
# Edit .env as needed (set REST_PORT, DB_DEBUG, DSN)
docker-compose up -d
```

### Manual Setup (Go)
```sh
git clone https://github.com/patricksferraz/loyalty-card.git
cd loyalty-card
cp .env.example .env
# Edit .env as needed
go mod tidy
go run main.go
```

### Environment Variables
See `.env.example`:
- `REST_PORT` - Port for REST API
- `DB_DEBUG` - Enable DB debug logs
- `DSN` - Database connection string

---

## Usage

### API Endpoints
All endpoints are prefixed with `/api/v1`.

#### Guests
- `POST   /guests` - Create a new guest
- `GET    /guests/{guest_id}` - Retrieve guest by ID

#### Scores
- `POST   /scores` - Create a new score
- `GET    /scores/{score_id}` - Retrieve score by ID
- `POST   /scores/{score_id}/use` - Use a score
- `POST   /scores/{score_id}/tags` - Add a tag to a score

#### Tags
- `POST   /tags` - Create a new tag
- `GET    /tags` - Search tags
- `GET    /tags/{tag_id}` - Retrieve tag by ID

### Example: Create a Guest
```sh
curl -X POST http://localhost:<REST_PORT>/api/v1/guests \
  -H 'Content-Type: application/json' \
  -d '{"name": "John Doe", "doc": "123456789"}'
```

### API Documentation
- [Swagger YAML](app/rest/docs/swagger.yaml)
- [Swagger JSON](app/rest/docs/swagger.json)
- Live Swagger UI: `http://localhost:<REST_PORT>/api/v1/swagger/index.html`

---

## Development

### Makefile Commands
- `make build`   - Build Docker image
- `make up`      - Start services (Docker Compose)
- `make down`    - Stop and remove containers
- `make logs`    - Tail logs
- `make gtest`   - Run Go tests with coverage
- `make attach`  - Open shell in running container

### Running Tests
```sh
make gtest
```

### Contributing
Contributions are welcome! Please open issues or submit pull requests. For major changes, open a discussion first.

---

## Tech Stack
- **Language:** Go 1.18+
- **Web Framework:** [Fiber](https://gofiber.io/)
- **ORM:** [GORM](https://gorm.io/)
- **CLI:** [Cobra](https://github.com/spf13/cobra)
- **API Docs:** [Swagger/OpenAPI](https://swagger.io/)
- **Containerization:** Docker, Docker Compose
- **Database:** PostgreSQL, SQLite

---

## License

This project is licensed under the [GNU GPLv3](LICENSE).

---

## Contact & Community
- **Author:** Patrick Ferraz
- **Email:** patrick.ferraz@coding4u.tech
- **GitHub Issues:** [Report bugs or request features](https://github.com/patricksferraz/loyalty-card/issues)
- **Discussions:** [Join the conversation](https://github.com/patricksferraz/loyalty-card/discussions)

---

## Acknowledgements
- [Fiber](https://gofiber.io/)
- [GORM](https://gorm.io/)
- [Cobra](https://github.com/spf13/cobra)
- [Swagger](https://swagger.io/)
- [Docker](https://www.docker.com/)
- Open-source community
