package db

import (
	"fmt"
	"log"
	"os"

	"github.com/ijaybaihaqi/heli-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct{}

func (p *Postgres) Connect(env string) (*gorm.DB, error) {
	var dbStr string

	if env == "development" {
		creds := model.Credential{
			Host:         "localhost",
			Username:     "postgres",
			Password:     "postgres",
			DatabaseName: "heli_db",
			Port:         5432,
		}

		dbStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)
	} else {
		dbUrl, ok := os.LookupEnv("DATABASE_URL")
		if !ok {
			log.Fatal("DATABASE_URL is must be set")
		}
		dbStr = dbUrl
	}

	dbConn, err := gorm.Open(postgres.Open(dbStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func NewDB() *Postgres {
	return &Postgres{}
}
