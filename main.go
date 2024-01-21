package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/freightcms/carriers/api"
	"github.com/freightcms/carriers/mongodb"

	"github.com/joho/godotenv"
)

type appEnvironment struct {
	mongodbURL string
	port       int
	host       string
}

func getAppEnv() *appEnvironment {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	return &appEnvironment{
		mongodbURL: os.Getenv("MONGO_URI"),
		port:       port,
		host:       os.Getenv("HOST"),
	}
}

func main() {

	env := getAppEnv()
	db, err := mongodb.NewCarrierDb(env.mongodbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create a new carrier service.
	app := api.CreateApp(db)
	// CreateApp creates an echo app with all routes defined.
	host := fmt.Sprintf("%s:%d", env.host, env.port)
	if err := app.Start(host); err != nil {
		panic(err)
	}
}
