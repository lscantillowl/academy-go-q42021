package main

import (
	"go_bootcamp_api/infrastructure/router"
	//good practices convention internals/personal/external
)

func main() {
	a := router.App{}
	a.Initialize()
	a.Run(":8080")
}
