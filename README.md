# Carriers

## Summary

Web API (Service) that provides capabilities for managing carrier information and related data such as drivers and compliance information.

## Development

### Setup

#### MongoDB
install mongodb and run https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-debian/

`sudo systemctl start mongod` starts the server
`sudo systemctl daemon-relate` restarts and reinitializes the service

Additional Commands can be found at https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-debian/#start-mongodb

#### Go work

```txt
go 1.23.4

use (
	.
	./db
	./db/mongodb
	./models
	./web
)
```

#### .env file

```dotenv
MONGO_SERVER=mongodb://root:example@0.0.0.0:27017/
```

### Running

From the root of the application...

1. Start Docker mongo database instance, and mongo express instance

  - If you would rather run mongodb on your local machine instead of installing docker please see [install-mongodb-on-debian](https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-debian/)
  - Creating the user
    1. Start the mongosh client by typing `mongosh`
    2. connect to the admin database with `use admin`
    3. Create a new user with `db.createUser({user: "root", pwd: "example", roles: [{role: "userAdminAnyDatabase", db: "admin"},{role: "readWriteAnyDatabase", db: "admin"}]})`
    4. Switch to the freightcms db with `use freightmcs`
    5. Create a new user with `db.createUser({user: "root", pwd: "example", roles: [{role: "userAdminAnyDatabase", db: "admin"},{role: "readWriteAnyDatabase", db: "admin"}]})`

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
