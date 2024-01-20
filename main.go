package main

import (
	"fmt"

	"github.com/freightcms/carriers/api"
	"github.com/freightcms/carriers/mongodb"
	"github.com/freightcms/carriers/services"
)

func main() {
	db, err := mongodb.NewCarrierDb("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create a new carrier service.
	service := services.NewCarrierService(db)
	app := api.CreateApp(service)
	// CreateApp creates an echo app with all routes defined.
	fmt.Print(app.Start(":8080"))
}
