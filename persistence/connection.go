package persistence

import (
	"fmt"
	"sync"

	"github.com/goakshit/gandalf/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Singleton - Connection pooling is handled
var db *gorm.DB
var once sync.Once

func getGORMConfig() *gorm.Config {
	return &gorm.Config{
		// Ignore default transaction started by GORM. Improves performance by upto 30%
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// Doesn't pluralize the table names
			// Eg: 'user' table won't be pluralized to 'users' table
			SingularTable: true,
		},
	}
}

// Return postgres connection string
func getPostgresConnString() string {
	config := config.New()
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.Database.User, config.Database.Passwd, config.Database.Host, config.Database.Port, config.Database.DatabaseName)
}

// GetGormClient - Returns db client
func GetGormClient() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(getPostgresConnString()), getGORMConfig())
		if err != nil {
			panic("Failed to open postgres connection\n" + err.Error())
		}
	})

	return db
}
