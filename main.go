package main

import (
	"github.com/lscantillowl/academy-go-q42021/infrastructure/router"
	//good practices convention internals/personal/external
)

func main() {
	a := router.App{}
	a.Initialize()
	a.Run(":8080")
}
