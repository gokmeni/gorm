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

func GetConnection() *gorm.DB {
	once.Do(func() {
		var err error

		db, err = gorm.Open("postgres", os.Getenv("POSTGRES_CONNECTION_STRING"))

		if err != nil {
			panic(err)
		}

		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.DB().SetConnMaxLifetime(time.Minute)

		err = db.DB().Ping()

		if err != nil {
			panic(err)
		}
	})

	return db
}
