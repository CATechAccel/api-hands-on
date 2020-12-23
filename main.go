package main

import (
	"github.com/fumist23/api-handson/router"
)

func main() {
	e := router.Router()
	e.Logger.Fatal(e.Start(":8080"))
}
