package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "sync"
    "workplace/internal/config"
)

var (
    connection *gorm.DB
    once       sync.Once
)

func GetConnection(configuration *config.Configuration) (db *gorm.DB, err error) {
    once.Do(func() {
        connection, err = gorm.Open(mysql.Open(configuration.DbDsn), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Silent),
        })
    })
    
    return connection, err
}
