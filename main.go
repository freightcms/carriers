package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	dotenv "github.com/dotenv-org/godotenvvault"
	"github.com/freightcms/carriers/db/mongodb"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func addMongoDbMiddleware(client mongo.Client, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := client.StartSession()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			w.Header().Set("ContentType", "application/json")
			return
		}
		defer session.EndSession(r.Context())

		sessionContext := mongo.NewSessionContext(r.Context(), session)
		personManagerContext := mongodb.WithContext(sessionContext)
		ctx := r.WithContext(personManagerContext)
		h.ServeHTTP(w, ctx)
	})
}

var (
	port int
	host string
)

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

	server.GET("/", echo.HandlerFunc(func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, &struct {
			Status string `json:"status" xml:"status"`
		}{
			Status: "Ok",
		}, "    ")
	}))
	addMongoDbMiddleware(server)
	fmt.Println("Done")
	hostname := fmt.Sprintf("%v:%d", host, port)
	fmt.Printf("Start server at %s\n", hostname)
	http.ListenAndServe(hostname, server)
}
