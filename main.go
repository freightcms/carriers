package main

import (
	"context"
	"log"
	"os"

	"github.com/freightcms/carriers/api"
	"github.com/freightcms/carriers/db/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	connStr := os.Getenv("CONNECTION_STRING")
	l := log.New(os.Stdin, "carriers: ", log.LstdFlags)
	serverApiVersion := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connStr).SetServerAPIOptions(serverApiVersion)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, struct {
			Hello string
		}{
			Hello: "world",
		})
	})

	r.Use(func(ctx *gin.Context) {
		session, err := client.StartSession()
		if err != nil {
			panic(err)
		}
		defer session.EndSession(ctx.Request.Context())
		appDb := mongodb.CreateCarrierDb(session)
		ctx.Set("db", appDb)
		ctx.Next()
	})
	api.CreateRouterGroup(r)
	l.Println("Listening on port 3000")
	l.Printf("database connnection %s", connStr)
	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
