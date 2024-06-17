# Carriers

Collection of schemas, models, and services for Managing Freight Carriers

## Workspaces

- api
- db
- db/mongodb
- models
- validators

# Testing

## Golang Terminal

run using `go test`

## Golang coverage

Run the below commands from the package you want to validate code coverage

`go test -v -coverprofile cover.out .` // where the ending "." can also be replaced from ./api, ./db
`go tool cover -html=cover.out`

## TODO