package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()
	e.Static("/swagger-ui.css", "dist/swagger-ui.css")
	e.Static("/swagger-ui-bundle.js", "dist/swagger-ui-bundle.js")
	e.Static("/swagger-ui-standalone-preset.js", "dist/swagger-ui-standalone-preset.js")
	e.Static("/swagger.yaml", "swagger.yaml")
	e.Static("/swaggerui", "dist/index.html")
	e.Logger.Fatal(e.Start(":8080"))
}
