package main

import (
	"fmt"

	"github.com/aruizgg/poke-api/app.go"
)

const port = ":3000"

func main() {
	router := app.SetupRouter()
	fmt.Printf("Server started at port %s\n", port)
	router.Run(port)
}
