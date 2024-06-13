
# Restaurant reservation system(Payment Microservice)

The Payment Microservice handles all payment-related functionalities for the restaurant reservation system. This includes processing payments, handling refunds, and managing payment methods.

## Overview

This microservice provides APIs to:
- **Process Payments**: Handle the payment process when a reservation is made.
- **Handle Refunds**: Manage refunds for canceled reservations.
- **Manage Payment Methods**: Allow users to add, update, or delete their payment methods.

## Technologies

- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Framework**: GIN
- **External Services**: Stripe for payment processing

## Installation

To set up the Payment Microservice, follow these steps:

### Prerequisites

- Go (version 1.16+)
- PostgreSQL
- Stripe API keys

### Setup

### **1. Clone the Repository**

```bash
  git init
  git clone git@github.com:dilshodforever/payment-service.git
```

### 2. Install Dependencies

```
go mod tidy
```


### 3. Update the .env file with the appropriate configuration.

```.env
    HTTPPort=:7070
    PostgresHost=localhost
    PostgresPort=5432
    PostgresUser=postgres
    PostgresPassword=root
    PostgresDatabase=reservation
```

### 4. Use the following Makefile commands to manage the database migrations and  set up the project:

```makefile
CURRENT_DIR=$(shell pwd)
DBURL=postgres://postgres:root@localhost:5432/newdb?sslmode=disable

SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./api/api.go -o $(SWAGGER_DOCS)

swag-gen:
  $(SWAGGER_INIT)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path migrations -database '$(DBURL)' -verbose up

mig-down:
	migrate -path migrations -database '$(DBURL)' -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table

```

### 5. Open the following URL to access the Swagger documentation:

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
