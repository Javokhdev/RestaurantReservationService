
# Restaurant reservation system(Api-Gateway)

The API Gateway Microservice acts as a single entry point for the restaurant reservation system, aggregating and routing requests to the appropriate microservices (e.g., reservation, payment, authorization). It handles request validation, response aggregation, rate limiting, and security.

## Overview

This microservice provides APIs to:
- **Route Requests**: Forward client requests to the appropriate microservices.
- **Aggregate Responses**: Combine responses from multiple microservices into a single response.
- **Handle Security**: Implement authentication and authorization checks.
- **Manage Rate Limiting**: Control the rate of incoming requests to prevent abuse.

## Technologies

- **Language**: Go (Golang)
- **API**: RESTful
- **Middleware**: Gorilla Mux, JWT Middleware
- **API Gateway Framework**: Kong, NGINX, or custom-built

## Installation

To set up the API Gateway Microservice, follow these steps:

### Prerequisites

- Go (version 1.16+)
- Kong or NGINX if using an existing API Gateway framework

### Setup

### 1. Clone the Repository

```bash
    git init 
    git clone git@github.com:dilshodforever/restaurant-api-gatway.git
```


### 2. Install Dependencies

```
go mod tidy
```


### 3. Configure Environment Variables.

```.env
AUTH_SERVICE_URL=http://localhost:8081
RESERVATION_SERVICE_URL=http://localhost:8082
PAYMENT_SERVICE_URL=http://localhost:8083
JWT_SECRET=your_jwt_secret_key

```

### 4. Start the Microservice

```
go run cmd/server/main.go
```


### 5. Use the following Makefile commands to manage the database migrations and  set up the project:

```makefile
CURRENT_DIR=$(shell pwd)
DBURL=postgres://postgres:root@localhost:5432/newdb?sslmode=disable

SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./api/api.go -o $(SWAGGER_DOCS)

swag-gen:
  $(SWAGGER_INIT)
```

### 6. Open the following URL to access the Swagger documentation:

```
http://localhost:8080/api/swagger/index.html#/
```



## Acknowledgements

#### Dilshod Umarov
[![telegram](https://img.shields.io/badge/telegram-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://t.me/DiLwOd_FoReVeR)

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/javohir-xasanov/)

#### Javoxir Xasanov 
[![telegram](https://img.shields.io/badge/telegram-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://t.me/javohir_khasanov)

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/javohir-xasanov/)

#### Muhammadsodiq Xudoynazarov 

[![telegram](https://img.shields.io/badge/telegram-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://t.me/XM_Mukhammed)

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/muhammadjon-xudaynazarov-89894b294)

#### Najmiddin Solihov
[![telegram](https://img.shields.io/badge/telegram-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://t.me/Salikhov079)

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/najmiddin-solihov-ab09612b2/)
