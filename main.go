package main

import (
	"github.com/lscantillowl/academy-go-q42021/config"
	"github.com/lscantillowl/academy-go-q42021/infrastructure/router"
)

func main() {
	a := router.App{}
	a.Initialize()
	a.Run(config.Port)
}
