package main

import (
	"project-alta-store/config"
	"project-alta-store/routes"
)

func main() {
	config.InitDB()
	e := routes.Start()
	e.Logger.Fatal(e.Start(":8000"))
}
