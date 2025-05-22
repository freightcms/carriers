package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	dotenv "github.com/dotenv-org/godotenvvault"
	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/db/mongodb"
	"github.com/freightcms/carriers/web"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// addMongoDbMiddleware adds the CarrierResourceManager to the echo context so that it can be
// be recovered from the db.DbContext object
func addMongoDbMiddleware(client *mongo.Client, next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		session, err := client.StartSession()
		if err != nil {
			return err
		}
		requestContext := c.Request().Context()
		defer session.EndSession(requestContext)

		sessionContext := mongo.NewSessionContext(requestContext, session)
		dbContext := db.DbContext{
			Context:                requestContext,
			CarrierResourceManager: mongodb.NewCarrierManager(sessionContext),
		}
		wrappedContext := web.AppContext{
			Context:   c,
			DbContext: dbContext,
		}
		return next(wrappedContext)
	})
}

var (
	port int
	host string
)

func dbContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Code to execute before the handler
		fmt.Println("Before handler")

		// Call the next handler
		err := next(c)

		// Code to execute after the handler
		fmt.Println("After handler")
		return err
	}
}

func main() {

	flag.IntVar(&port, "p", 8080, "Port to run application on")
	flag.StringVar(&host, "h", "0.0.0.0", "Host address to run application on")
	ctx := context.Background()
	fmt.Println("Starting application...")

	if err := dotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_SERVER")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer client.Disconnect(ctx)
	fmt.Println("Pinging server...")
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Done")
	fmt.Println("Setting up handlers and routes")

	server := echo.New()
	server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return addMongoDbMiddleware(client, next)
	})

	web.Register(server)

	server.GET("/", echo.HandlerFunc(func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, &struct {
			Status string `json:"status" xml:"status"`
		}{
			Status: "Ok",
		}, "    ")
	}))
	fmt.Println("Done")
	hostname := fmt.Sprintf("%v:%d", host, port)
	fmt.Printf("Start server at %s\n", hostname)
	http.ListenAndServe(hostname, server)
}
