package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "sync"
)

type Configuration struct {
    DbDsn              string `yaml:"database_dsn"`
    GrpcHost           string `yaml:"grpc_host"`
    GrpcPort           string `yaml:"grpc_port"`
    RedisUrl           string `yaml:"redis_url"`
    LogPath            string `yaml:"log_path"`
    RotateLogPath      string `yaml:"rotate_log_path"`
    HttpHost           string `yaml:"http_host"`
    HttpPort           string `yaml:"http_port"`
    ApiToken           string `yaml:"api_token"`
    ApiBillingUrl      string `yaml:"api_billing_url"`
    ApiBillingUser     string `yaml:"api_billing_user"`
    ApiBillingPassword string `yaml:"api_billing_password"`
    TelegramLogBotKey  string `yaml:"api_telegram_log_bot_key"`
    TelegramLogBotChat string `yaml:"api_telegram_log_bot_chat_id"`
}

var (
    configuration *Configuration
    once          sync.Once
)

func GetConfig() *Configuration {
    once.Do(func() {
        file, err := ioutil.ReadFile("config.yaml")
        if err != nil {
            log.Fatal("cannot read configuration:", err)
        }
        configuration = &Configuration{}
        if err = yaml.Unmarshal(file, configuration); err != nil {
            log.Fatal("cannot unmarshal config.yaml", err.Error())
        }
    })
    
    return configuration
}
