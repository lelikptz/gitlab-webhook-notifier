package notifier

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
)

type TelegramNotifier struct {
}

func NewTelegramNotifier() *TelegramNotifier {
	return &TelegramNotifier{}
}

func (t *TelegramNotifier) Send(message string) {
	bot, err := telegram.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Printf("something went wrong %v", err)
		return
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Printf("something went wrong %v", err)
		return
	}
	bot.Debug = debug

	chatID, err := strconv.Atoi(os.Getenv("TELEGRAM_CHAT_ID"))
	if err != nil {
		log.Printf("something went wrong %v", err)
		return
	}

	msg := telegram.NewMessage(int64(chatID), message)
	_, err = bot.Send(msg)
	if err != nil {
		log.Printf("something went wrong %v", err)
	}
}
