package connection

import (
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {

	dsn := os.Getenv("DATABASE_URL")

	maxTries := 10
	for i := 0; i < maxTries; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("success connect to mysql")
			return db, nil
		}

		log.Println("failed to connect to mysql, try again in 1 minute")
		time.Sleep(1 * time.Minute)
	}

	return nil, errors.New("mysql connection failed after 10 minute")
}
