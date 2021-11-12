package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"workplace/config"
)

var (
	connection *gorm.DB
	once       sync.Once
)

func GetConnection() (db *gorm.DB, err error) {
	once.Do(func() {
		configuration := config.GetConfig()
		connection, err = gorm.Open(mysql.Open(configuration.DbDsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	})

	return connection, err
}
