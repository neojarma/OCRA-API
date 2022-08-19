package test

import (
	"context"
	"ocra_server/router"
	firebase_service "ocra_server/service/firebase"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type SetupModel struct {
	App             *echo.Echo
	Dialer          *gomail.Dialer
	Db              *gorm.DB
	Group           *echo.Group
	FirebaseService firebase_service.FirebaseService
}

func Setup() *SetupModel {
	app := echo.New()

	// setup db connection
	db, err := GetConnection()
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

	group := app.Group("/api/v1")

	router.Router(group, db, dialer, firebaseService)

	return &SetupModel{
		App:             app,
		Dialer:          dialer,
		Db:              db,
		Group:           group,
		FirebaseService: firebaseService,
	}
}
