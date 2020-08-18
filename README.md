# prods-api
`prods-api` is a go based API server that supports basic CRUD operations for categories and products. Each product can belong up to a single category. 

# Getting Started

## Prerequisites
You will need to have Go v1.12 installed and configured in your environment. Also, docker and docker-compose can be used to provision a MySQL DB that will be used from app and a redis server used for caching GET endpoints. If docker and docker-compose is not available other MySQL an Redis instances can be used as well. Redis is not a mandatory dependency for the project to run.

DDL and sample data scripts are locates under the folder `api/repositories/`.

## Setting Up
In order to start a MySQL server and a Redis service by running run under project's root directory:
```
docker-compose up
```
Then start `prods-api` app by running under project's root directory
```
go run main
```

If everything works normally this will output:
```
Listening at :8080
```

# Usage

## Configuration
Application configuration is in the `.env' file. They can be store in this file or in ENV_VARS of the OS.
`SERVER_PORT` is the port that will be used by the Web server of this project and `API_PREFIX` is the default prefix for all API endpoints
Fields prefixed with `MYSQL_` provide details for connecting to the MySQL server which will be used for the application. The credentials are also used within the `docker-compose.yml` file to instantiate the DB. If `MIGRATE_DB` is set to true, the DB schema will be re-generated in the MySQL and if `SEED_DATA` is set to true all DB's data will be truncated and some sample data will be inserted.

```
# Server
SERVER_PORT=8080
API_PREFIX=/api

# MySQL
MYSQL_HOST=127.0.0.1
MYSQL_DATABASE=prods_db
MYSQL_USER=prods_db_user
MYSQL_PASSWORD=prods_db_password
MYSQL_ROOT_PASSWORD=password

# Repository
MIGRATE_DB=false
SEED_DATA=false
```

## API Reference
In order to review the provided API a working swaggerUI is set up with this app and runs at this link:
```
http://localhost:8080/swagger/index.html
```
while the OpenAPI spec is served at this endpoint:
```
http://localhost:8080/api-docs.json
```

Errors that occur in the API respond with the following structure:
```
{
    "trace_id": "969fbf24-03bb-4b3c-a795-12859d37da25",
    "timestamp": "2020-05-25T20:06:40+03:00",
    "message": "Invalid Category.",
    "code": "invalid",
    "http_status_code": 400,
    "http_status": "Bad Request"
}
```

# Functionality
API provides basic CRUD operation on `products` and `categories`. The used models are the following:
```
Product:
{
    id	integer
    *title	string
    *price	integer
    category_id	integer
    description	string
    image_url	string
    created_at	string
    updated_at	string
}

Category:
{
    id	integer
    *title	string
    image_url	string
    sort	integer
    created_at	string
    updated_at	string
}
```
(*) Mandatory fields for POST/PUT

Products' price is manipulated as price in CENTS of the currency from the DB up to the API.

# Tests
in order to run the available Unit Tests run:
```
go test ./...
```
Note: Tests are written in various ways to demonstrate the capabilities of different approaches. Not all cases have covered.

# Contributing Guidelines
This application has developed using a layered architectural style. Supported layers are repository, service and controller. Each layer has separate and distinct role and is dependant only in the layer bellow its level so that we can maintain and expand it's functionality.

Controllers and repositories are using the models within the repository layer, while controllers
are communicating with the services through these models and convert to DTOs to serve the API's end-users.

In any change of the API contract please update the OpenAPI docs by running this command:
```
swag init -g main.go --output docs
```

# Future Improvements
* Add integration tests at the repository layer towards a test DB. Integration test are crucial at this level as logic is enforced through the DB and also querying of data is only being done through the DB.
* Add integration tests at the controller layer. Integration test are crucial at this level as we can test among the API contract that our end-users use.
* Use migration scripts and logic to track DB's schema updates and rollbacks.
* Exploit cache-control headers of the requests to manipulate caching and expiration of data.
