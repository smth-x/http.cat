package main

import (
	"fmt"
	"http.cat/src/const"
	"http.cat/src/telegram-bot"
	ut "http.cat/src/utils"
)

func main() {
	runBot()
}
func runBot() {
	err := telegram_bot.CreateBot(_const.BotToken)
	if !ut.IsErrorNil(err) {
		fmt.Println(err)
		return
	}
	telegram_bot.StartLongPolling(telegram_bot.Bot)
}
