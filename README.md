# Carriers

Stand alone Application containing various libraries that can be used to stand up an API which can hold the data and inforation for Freight Carriers. The different Packages can be used independantly of each other through Inversion of Control (IoC) and Dependency Injection Practices.

## Tech Stack

- Golang - language built in
- MongoDB - database
- Echo - Web API Framework

## Getting Started

### Prerequisites

- Have [docker](https://www.docker.com/products/docker-desktop/) installed
- Have [mongo shell](https://www.mongodb.com/docs/mongodb-shell/install/) installed
- Have the version of Go pinned in the `go.mode` file

### Environment Setup

1. Run `docker compose pull` to get latest
2. Clone your respository

```bash
`git clone https://github.com/freightcms/carriers`
```

3. Move into the cloned directory

```bash
`cd ./carriers`
```

4. Create an `.env` file in the root fo the directory

```bash
touch .env
```

```text
HOST=0.0.0.0
PORT=5000
DEBUG=True
LOG_LEVEL=DEBUG
MONGO_URI=mongodb://localhost:27017
```

5. Fetch `mongodb` image

```sh
docker compose pull
```

6. Start `mongodb` instance

```sh
docker compose up mongodb
```

7. Seed the database and create collections by running script in terminal `./scripts/init_db.sh`
   - If you run into an issue with execution run `chmod +x ./scrpts/*`
8. make sure you have all the dependencies installed

```sh
go mod tidy
```

6. run the project

```sh
go run main.go
```

## Projects

### api

Contains all the resources necessary to stand up an API and handle basic CRUD operations.

### models

Database models and optionally API schema models that can be used for requests.

### db

Contains an `interface` for the APi/Service/mongodb libraries to share and exchange information.

### mongodb

Contains logic for reading and writing to the carriers mongodb database.

### services

Contains business logic to exchange information between API and Database layers.

### Schemas

Contains information for APIs to serialize and deserialize struct objects to/from `JSON` and `BSON`.

## See Also

- [Freight KB Carriers](https://kb.freightcms.com/carriers/)
- [mongodb driver](https://github.com/mongodb/mongo-go-driver)
- [mongoshell](https://www.mongodb.com/docs/mongocli/stable/command/mongocli/)
- [godot](https://golangbyexample.com/load-env-fiie-golang/)