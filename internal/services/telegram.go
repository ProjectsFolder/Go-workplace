package services

import (
    "fmt"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "log"
    "strconv"
    "time"
    "workplace/internal/config"
    stringUtils "workplace/internal/utils"
)

const TelegramMessageMaxLength = 4096

type Telegram struct {
    bot *tgbotapi.BotAPI
    chatId int64
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

func (t *Telegram) Log(v ...string) {
    go func() {
        now := time.Now()
        timeFormatted := now.Format("02.01.2006 15:04:05")
        message := fmt.Sprintf("[%s]: ", timeFormatted)
        for _, arg := range v {
            message += " " + arg
        }

        chunks := stringUtils.StringChunk(message, TelegramMessageMaxLength)
        for _, chunk := range chunks {
            msg := tgbotapi.NewMessage(t.chatId, chunk)
            t.bot.Send(msg)
        }
    }()
}
