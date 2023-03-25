package main

import (
	"os"

	"github.com/sk25469/momoney-backend-assignment/routes"
)

func main() {
	port, available := os.LookupEnv("PORT")
	if !available {
		port = "3000"
	}
	routes.RegisterRoutes(port)
}
