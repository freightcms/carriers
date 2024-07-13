# Carriers

Collection of schemas, models, and services for Managing Freight Carriers

## Workspaces

- api
- db
   - mongodb
- models
- schemas

## Running Locally

1. Install all dependencies

```sh
go get && go build
```

2. Install Docker
3. Start mongodb and mongodb-express

```sh
docker compose up -d mongo mongo-express
```

4. Start your local application

```sh
go run ./main.go
```


# Testing

## Golang Terminal

run using `go test`

## Golang coverage

Run the below commands from the package you want to validate code coverage

`go test -v -coverprofile cover.out .` // where the ending "." can also be replaced from ./api, ./db
`go tool cover -html=cover.out`

## TODO

- Basic CRUD Operations for Carriers
  - models are created
  - schemas are created
  - validatoin is created for schemas
  - models have their database properties set, such as types, widths of column size

