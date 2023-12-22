package main

import (
	"fmt"

	"github.com/freightcms/carriers/api"
)

func main() {
	app := api.CreateApp()
	// CreateApp creates an echo app with all routes defined.
	fmt.Print(app.Start(":8080"))
}
