package main

import (
	"context"
	"ocra_server/connection"
	"ocra_server/router"
	firebase_service "ocra_server/service/firebase"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/gomail.v2"
)

var setupService *router.SetupService

func init() {
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

	// setup dialer
	dialer := &gomail.Dialer{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     port,
		Username: os.Getenv("SMTP_EMAIL"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}

	// setup firebase
	firebaseConfigClient := firebase_service.GetFirebaseStorageClient(context.Background())
	firebaseService := firebase_service.NewFirebaseService(firebaseConfigClient)

	setupService = &router.SetupService{
		Db:              db,
		Dialer:          dialer,
		FirebaseService: firebaseService,
	}
}

func main() {
	app := echo.New()
	setupService.Group = app.Group("/api/v1")

	setupService.Group.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"https://web.ocra.neojarma.com"},
	}))

	router.Router(setupService)

	app.Start(":8080")
}
