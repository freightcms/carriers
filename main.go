package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	dotenv "github.com/dotenv-org/godotenvvault"
	"github.com/graphql-go/handler"
	"github.com/squishedfox/webservice-prototype/db/mongodb"
	"github.com/squishedfox/webservice-prototype/web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	rootSchema, err := web.NewSchema()
	if err != nil {
		log.Fatal(err)
		return
	}
	h := handler.New(&handler.Config{
		Schema:   &rootSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	server := http.NewServeMux()
	server.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
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

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":\"ok\"}"))
		w.Header().Set("ContentType", "application/json")
	})
	fmt.Println("Done")
	hostname := fmt.Sprintf("%v:%d", host, port)
	fmt.Printf("Start server at %s", hostname)
	http.ListenAndServe(hostname, server)
}
