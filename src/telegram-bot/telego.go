package telegram_bot

import (
	"fmt"
	"http.cat/src/const"
	rq "http.cat/src/request"
	ut "http.cat/src/utils"
	"os"
	"regexp"
	"strconv"
	"strings"

	tg "github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var Bot *tg.Bot
var ID int64

func CreateBot(token string) error {
	createdBot, err := tg.NewBot(token, tg.WithDefaultDebugLogger())
	if !ut.IsErrorNil(err) {
		return err
	}

	Bot = createdBot
	return nil
}
func StartLongPolling(bot *tg.Bot) {
	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		err := handle(update)
		if !ut.IsErrorNil(err) {
			fmt.Println(err)
		}
	}
}
func handle(update tg.Update) error {
	ID = update.Message.Chat.ID
	messageText := strings.ToLower(update.Message.Text)

	if strings.HasPrefix(messageText, "/get ") {
		re := regexp.MustCompile(`get (\d+)`)
		matches := re.FindStringSubmatch(messageText)

		//getting the http code
		if len(matches) > 1 {
			number, err := strconv.Atoi(matches[1])
			if !ut.IsErrorNil(err) {
				return err
			}

			//get image from API
			err = rq.GetRequest(_const.FileName, number)
			if !ut.IsErrorNil(err) {
				return err
			}
			defer rq.DeleteImg()

			return sendPhoto(_const.CaptionToImage, number)
		}
	} else if messageText == "/start" {
		return sendMessage(_const.StartMessage)
	} else if messageText == "/commandlist" {
		return sendMessage(getCommandList())
	} else if messageText == "/get" {
		return sendMessage(_const.EmptyGetCommand)
	} else {
		return sendMessage(_const.UnknownMessage)
	}
	return nil
}
func getCommandList() (out string) {
	if len(_const.Commands) == 0 {
		return ""
	}
	out += fmt.Sprintf("- Command list -\n")
	for i := 0; i < len(_const.Commands); i++ {
		currentCommand := _const.Commands[i]

		out += fmt.Sprintf("· %s - %s", currentCommand.Command, currentCommand.Description)
		if _const.Commands[i].ExampleOfUsing != "" {
			out += fmt.Sprint("\n    exp: ", currentCommand.ExampleOfUsing)
		}

		//if current command isn`t last in list
		if len(_const.Commands)-i != 1 {
			out += fmt.Sprintf("\n")
		}
	}
	return
}
func sendMessage(message string) error {
	_, err := Bot.SendMessage(tu.Message(tg.ChatID{ID: ID}, message))
	return err
}
func sendPhoto(caption string, httpCode int) error {
	file, err := os.Open(_const.FileName)
	if !ut.IsErrorNil(err) {
		return err
	}
	defer file.Close()

	photo := tu.Photo(
		tg.ChatID{ID: ID},
		tu.File(file),
	).WithCaption(fmt.Sprintf("%s%d", caption, httpCode))

	_, _ = Bot.SendPhoto(photo)
	return nil
}
