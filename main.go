package main

import (
	"context"
	"os"

	carrierApi "github.com/freightcms/carriers/api"
	"github.com/freightcms/carriers/db/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	connStr := os.Getenv("CONNECTION_STRING")
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
	carrierApi.CreateRouterGroup(r)
	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
