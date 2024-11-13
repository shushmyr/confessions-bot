package handlers

import (
	"cf-bot/internal/keyboards"
	"cf-bot/internal/texts"
	"cf-bot/internal/users"
	"log"
	"strconv"

	tg "github.com/OvyFlash/telegram-bot-api"
)

func Start(chatID int64, bot *tg.BotAPI) {
  msg := tg.NewMessage(chatID, texts.Start)
  msg.ReplyMarkup = keyboards.StartKB

  bot.Send(msg) 
}

func TakeTxt(chatID int64, bot *tg.BotAPI) {
  msg := tg.NewMessage(chatID, texts.Take)
  msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
  msg.ReplyMarkup = keyboards.TakeKB

  bot.Send(msg)
}

func AnonTxt(chatID int64, bot *tg.BotAPI, adminsChatID, userID int64, usrMsg string, userName string) {
  userIDstr := strconv.Itoa(int(userID))
  msgToAdmins := tg.NewMessage(adminsChatID, usrMsg + "\n\n#тейк")
  msgToAdmins2 := tg.NewMessage(adminsChatID, "ID: " + userIDstr)
  bot.Send(msgToAdmins)
  bot.Send(msgToAdmins2)

  users.Add(chatID, userID)

  msg := tg.NewMessage(chatID, "тейк был отправлен админам")
  msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
  bot.Send(msg)

  log.Printf("анонимный тейк от @%s был отправлен в чат админов", userName)
}

func NeanonTxt(chatID int64, bot *tg.BotAPI, adminsChatID, userID int64, usrMsg string, userName string) {
  userIDstr := strconv.Itoa(int(userID))

  msgToAdmins := tg.NewMessage(adminsChatID, usrMsg + "\n\n#тейк")
  msgToAdmins2 := tg.NewMessage(adminsChatID, "тейк от @" + userName + " " + "ID: " + userIDstr)
  bot.Send(msgToAdmins)
  bot.Send(msgToAdmins2)

  users.Add(chatID, userID)

  msg := tg.NewMessage(chatID, "тейк был отправлен админам")
  msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
  bot.Send(msg)

  log.Printf("неанонимный тейк от @%s был отправлен в чат админов", userName)
}

func WontWriteTake(chatID int64, bot *tg.BotAPI) {
  msg := tg.NewMessage(chatID, "")
  msg.ReplyMarkup = tg.NewRemoveKeyboard(true) 
}

func AdmResponse(userID, adminsChatID int64, bot *tg.BotAPI, response, adminUserName string) {
  chatID := users.GetUser(userID)
  msg := tg.NewMessage(chatID, response)
  bot.Send(msg)

  msgToAdmins := tg.NewMessage(adminsChatID, "ответ отправлен")
  bot.Send(msgToAdmins)

  log.Printf("отправлен ответ пользователю %v от @%s", userID, adminUserName)
}
