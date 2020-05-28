package ds

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"sync"
	"time"
)

var db *gorm.DB
var once sync.Once

const (
	postgresEnvVariable string = "POSTGRES_CONNECTION_STRING"
)

func GetConnection() *gorm.DB {
	once.Do(func() {
		var err error

		db, err = gorm.Open("postgres", os.Getenv(postgresEnvVariable))

		if err != nil {
			panic(err)
		}

		db.DB().SetMaxIdleConns(20)
		db.DB().SetMaxOpenConns(20)
		db.DB().SetConnMaxLifetime(5 * time.Minute)

		err = db.DB().Ping()

		if err != nil {
			panic(err)
		}
	})

	return db
}
