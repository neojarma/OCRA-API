package main

import (
	"ocra_server/connection"
	"ocra_server/router"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/gomail.v2"
)

func main() {
	app := echo.New()

	// setup db connection
	db, err := connection.GetConnection()
	if err != nil {
		panic(err)
	}

	// setup mail service
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		panic(err)
	}

	dialer := &gomail.Dialer{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     port,
		Username: os.Getenv("SMTP_EMAIL"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}

	group := app.Group("/api/v1")

	group.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"https://web.ocra.neojarma.com"},
	}))

	router.Router(group, db, dialer)

	app.Start(":8080")
}
