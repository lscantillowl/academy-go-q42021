package main

import (
	"github.com/lscantillowl/academy-go-q42021/infrastructure/router"
)

// Main function to start the application and initialize the router
func main() {
	a := router.App{}
	a.Initialize()
	a.Run(":8080")
}
