# Carriers

Collection of Data Transfer Objects (DTOs) that represent the different types of Carriers and their related information.

## Getting Started

1. Ensure you have the version of Go pinned in the `go.mode` file
2. clone your respository

```bash
`git clone https://github.com/freightcms/carriers`
```

3. move into the cloned directory

```bash
`cd ./carriers`
```

4. Create an `.env` file in the root fo the directory

```bash
touch .env
```

```text
LOG_LEVEL=INFO
DEBUG=true
HOST=0.0.0.0
PORT=8080
```

4. make sure you have all the dependencies installed

```bash
go mod tidy
```

5. run the project

```bash
go run main.go
```

## Projects

### API

Contains all the resources necessary to stand up an API and handle basic CRUD operations.

### Models

Database models and optionally API schema models that can be used for requests.

## See Also

- [Freight KB Carriers](https://kb.freightcms.com/carriers/)
- [mongodb driver](https://github.com/mongodb/mongo-go-driver)