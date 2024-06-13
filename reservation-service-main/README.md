
# Restaurant reservation system(Reservation microservice)


This microservice handles the core functionalities related to managing restaurant reservations, 
including booking tables, selecting meals, processing payments, and managing reservations.

## Overview

The Reservation Microservice provides APIs to:

- **Check Table Availability**: Verify the availability of tables for a given date and time.
- **Book Tables**: Allow users to book available tables.
- **Select Meals**: Enable users to pre-select meals when making a reservation.


## Technologies

- **Language**: Go (Golang)
- **Protocol Buffers**: For data serialization and defining APIs
- **gRPC**: For internal service communication
- **Gin**: For external communication
- **Database**: PostgreSQL (or any other relational database)

To set up the Reservation Microservice, follow these steps:

### Prerequisites

- Go (version 1.16+)
- Protocol Buffers Compiler (`protoc`)
- PostgreSQL
- RabbitMQ (if using messaging for notifications)


**Clone Reservation Service:**
    
```bash
  git init
  git clone git@github.com:dilshodforever/reservation-service.git
```



### 2. Create a database named reservation on port: 5432


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
