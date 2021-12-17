package services

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "log"
    "strconv"
    "sync"
    "workplace/internal/config"
    stringUtils "workplace/internal/utils"
)

const TelegramMessageMaxLength = 4096

type Telegram struct {
    bot *tgbotapi.BotAPI
    chatId int64
    mutex sync.Mutex
}

func NewTelegramClient(configuration *config.Configuration) *Telegram {
    telegramBot, err := tgbotapi.NewBotAPI(configuration.TelegramLogBotKey)
    if err != nil {
        log.Fatal("Cannot create telegram-bot:", err)
    }
    chatId, err := strconv.ParseInt(configuration.TelegramLogBotChat, 10, 64)
    if err != nil {
        log.Fatal("Cannot create telegram-bot:", err)
    }
    return &Telegram{
        bot: telegramBot,
        chatId: chatId,
    }
}

func (t *Telegram) LogAsync(v ...string) {
    go func() {
        t.Log(v...)
    }()
}

func (t *Telegram) Log(v ...string) error {
    var message = ""
    for argNum, arg := range v {
        if argNum > 0 {
            message += " "
        }
        message += arg
    }
    
    chunks, _ := stringUtils.ChunkSplit(message, TelegramMessageMaxLength)
    for _, chunk := range chunks {
        msg := tgbotapi.NewMessage(t.chatId, chunk)
        _, err := t.bot.Send(msg)
        if err != nil {
            return err
        }
    }
    
    return nil
}

func (t *Telegram) Write(p []byte) (int, error) {
    t.mutex.Lock()
    defer t.mutex.Unlock()

    err := t.Log(string(p))
    if err != nil {
        return 0, err
    }

    return len(p), nil
}
