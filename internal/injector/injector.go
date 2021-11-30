package injector

import (
    "go.uber.org/dig"
    "gorm.io/gorm"
    "sync"
    "workplace/internal/config"
    "workplace/internal/database"
)

var (
    injector *dig.Container
    once     sync.Once
)

func buildContainer() *dig.Container {
    once.Do(func() {
        injector = dig.New()
    
        err := injector.Provide(config.GetConfig)
        if err != nil {
            panic(err)
        }
    
        err = injector.Provide(func(config *config.Configuration) (*gorm.DB, error) {
            db, err := database.GetConnection(config)
        
            return db, err
        })
        if err != nil {
            panic(err)
        }
    
        err = injector.Provide(func(db *gorm.DB) (*database.ProductRepositoryImpl, error) {
            return database.NewProductRepository(db), nil
        })
        if err != nil {
            panic(err)
        }
    })
    
    return injector
}

func GetContainer() *dig.Container {
    return buildContainer()
}
