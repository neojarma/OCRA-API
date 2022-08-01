package main

import (
	"ocra_server/connection"
	"ocra_server/router"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	db, err := connection.GetConnection()
	if err != nil {
		panic(err)
	}

	group := app.Group("/api/v1")
	router.Router(group, db)

	app.Start(":8080")
}
