# webservice-prototype

A prototype building out with GraphQL, Golang, and MongoDB

## Setup

### Go work

```txt
go 1.22.5

use (
    .
    ./db
    ./db/mongodb
    ./models
    ./web
)
```

### .env file

```dotenv
MONGO_SERVER=mongodb://root:example@0.0.0.0:27017/
```

## Running

From the root of the application...

1. Start Docker mongo database instance, and mongo express instance

```sh
docker compose up -d
```
2. Make sure to install all dependencies

```sh
go get
```

3. Start the main application

```sh
go run main.go
```

4. test the application is up and healthy

```sh
curl http://localhost:8080/
```

should output

```sh
{ "status": "ok" }
```
