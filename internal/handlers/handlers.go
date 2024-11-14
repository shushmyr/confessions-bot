package handlers

import (
	"cf-bot/internal/keyboards"
	"cf-bot/internal/texts"
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

func wontWriteTake(chatID int64, bot *tg.BotAPI) {
  msg := tg.NewMessage(chatID, "")
  msg.ReplyMarkup = tg.NewRemoveKeyboard(true) 
  bot.Send(msg)
} 

func AnonTxt(updates tg.UpdatesChannel, bot *tg.BotAPI, adminsChatID int64) {
  for u := range updates {
    //если сообщения нет пропускаем итерацию
    if u.Message == nil {
      continue
    }

    //переменные
    chatID := u.Message.Chat.ID
    msgText := u.Message.Text
    userName := u.SentFrom().UserName    
    userID := u.SentFrom().ID

  
    //если нажата кнопка отмены тейка
    if msgText == "не хочу отправлять тейк" {
      wontWriteTake(chatID, bot)

      //возврат в стартовое меню
      Start(chatID, bot)
      break
    }
    
    //отправка тейка
    userIDstr := strconv.Itoa(int(userID))

    msgToAdmins := tg.NewMessage(adminsChatID, msgText + "\n\n#тейк")
    msgToAdmins2 := tg.NewMessage(adminsChatID, "ID: " + "`" + userIDstr + "`")
    msgToAdmins2.ParseMode = "MarkdownV2"
    bot.Send(msgToAdmins)
    bot.Send(msgToAdmins2)

    msg := tg.NewMessage(chatID, "тейк был отправлен админам")
    msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
    bot.Send(msg)

    log.Printf("анонимный тейк от @%s был отправлен в чат админов", userName)
    //возврат в стартовое меню
    Start(chatID, bot)
    break
  }
}

func NeanonTxt(updates tg.UpdatesChannel, bot *tg.BotAPI, adminsChatID int64) {
  for u := range updates {
    //если сообщения нет пропускаем итерацию
    if u.Message == nil {
      continue
    }

    //переменные
    chatID := u.Message.Chat.ID
    msgText := u.Message.Text
    userName := u.SentFrom().UserName    
    userID := u.SentFrom().ID

  
    //если нажата кнопка отмены тейка
    if msgText == "не хочу отправлять тейк" {
      wontWriteTake(chatID, bot)

      //возврат в стартовое меню
      Start(chatID, bot)
      break
    }
    
    //отправка тейка
    userIDstr := strconv.Itoa(int(userID))

    msgToAdmins := tg.NewMessage(adminsChatID, msgText + "\n\n#тейк")
    msgToAdmins2 := tg.NewMessage(adminsChatID, "тейк от @" + userName + " ID: " + "`" + userIDstr + "`")
    msgToAdmins2.ParseMode = "MarkdownV2"
    bot.Send(msgToAdmins)
    bot.Send(msgToAdmins2)
    msg := tg.NewMessage(chatID, "тейк был отправлен админам")

    msg.ReplyMarkup = tg.NewRemoveKeyboard(true)
    bot.Send(msg)

    log.Printf("неанонимный тейк от @%s был отправлен в чат админов", userName)
    
    //возврат в стартовое меню
    Start(chatID, bot)
    break
  }
}
