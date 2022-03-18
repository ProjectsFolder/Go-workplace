package injector

import (
    "github.com/go-redis/redis"
    "go.uber.org/dig"
    "gorm.io/gorm"
    "log"
    "net/http"
    "sync"
    "time"
    "workplace/internal/config"
    "workplace/internal/database"
    httpJwt "workplace/internal/http/jwt"
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
            client := redis.NewClient(&redis.Options{
                Addr: config.RedisUrl,
                Password: "",
                DB: 0,
            })
            _, err = client.Ping().Result()

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

        err = injector.Provide(func(config *config.Configuration) (*services.Telegram, error) {
            return services.NewTelegramClient(config), nil
        })
        if err != nil {
            panic(err)
        }

        createLogger := func(telegram *services.Telegram) *log.Logger {
            return log.New(telegram, "", log.Ldate|log.Ltime|log.Lshortfile)
        }

        err = injector.Provide(func(telegram *services.Telegram) (*log.Logger, error) {
           return createLogger(telegram), nil
        })
        if err != nil {
           panic(err)
        }

        err = injector.Provide(func(telegram *services.Telegram) (*log.Logger, error) {
            return createLogger(telegram), nil
        }, dig.Name("telegramLogger"))
        if err != nil {
            panic(err)
        }

        err = injector.Provide(func(config *config.Configuration) (*httpJwt.Jwt, error) {
            return httpJwt.NewJWT(config.JwtSecret), nil
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
