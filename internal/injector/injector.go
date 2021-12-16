package injector

import (
    "github.com/go-redis/redis"
    "go.uber.org/dig"
    "gorm.io/gorm"
    "net/http"
    "sync"
    "time"
    "workplace/internal/config"
    "workplace/internal/database"
    "workplace/internal/services"
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

        err = injector.Provide(func(config *config.Configuration) (*redis.Client, error) {
            client, err := services.GetRedisClient(config)

            return client, err
        })
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

        err = injector.Provide(func(db *gorm.DB) (database.ProductRepository, error) {
            return database.NewProductRepository(db), nil
        })
        if err != nil {
            panic(err)
        }

        err = injector.Provide(func(config *config.Configuration) (*services.Logger, error) {
            return services.CreateLogger(config), nil
        })
        if err != nil {
            panic(err)
        }
    
        err = injector.Provide(func(config *config.Configuration) (services.BillingProviderInterface, error) {
            client := &http.Client{Timeout: 30 * time.Second}

            return services.NewBillingClient(
                client,
                config.ApiBillingUrl,
                config.ApiBillingUser,
                config.ApiBillingPassword,
            ), nil
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
