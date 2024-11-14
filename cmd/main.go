package main

import (
	"log"
	"os"

	"strconv"

	"cf-bot/internal/handlers"

	tg "github.com/OvyFlash/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("can't load token")
  }

  token := os.Getenv("TOKEN")
  adminsChatIDstr := os.Getenv("ADM_CHAT")
  adminsChatIDi, err := strconv.Atoi(adminsChatIDstr)
  if err != nil {
    log.Fatal("can't get admins chat id")
  }
  adminsChatID := int64(adminsChatIDi)

  bot, err := tg.NewBotAPI(token)

  log.Printf("авторизован под %s", bot.Self.UserName)

  updateConfig := tg.NewUpdate(0)
  updateConfig.Timeout = 60

  updates := bot.GetUpdatesChan(updateConfig)

  for u := range updates {
    //если нет сообщения пропускаем итерацию
    if u.Message == nil {
      continue
    }
    
    //переменные
    chatID := u.Message.Chat.ID
    msgText := u.Message.Text
    userName := u.SentFrom().UserName

    //обрабатываем текст
    if msgText != "" {
      //если от админов
      if chatID == adminsChatID {
         
      }
      //если не от админа
      switch msgText {
      //стартовое меню
      case "/start":
        logTxt(msgText, userName)
        handlers.Start(chatID, bot)
      case "анон":
        logTxt(msgText, userName)
        handlers.TakeTxt(chatID, bot)

        //цикл для ожидания тейка
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
            handlers.WontWriteTake(chatID, bot)
      
            //возврат в стартовое меню
            handlers.Start(chatID, bot)
            break
          }
          
          //отправка тейка
          handlers.AnonTxt(chatID, bot, adminsChatID, userID, msgText, userName)
          //возврат в стартовое меню
          handlers.Start(chatID, bot)
          break
        }
      case "неанон":
        logTxt(msgText, userName)
        handlers.TakeTxt(chatID, bot)

        //цикл для ожидания тейка
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
            handlers.WontWriteTake(chatID, bot)
      
            //возврат в стартовое меню
            handlers.Start(chatID, bot)
            break
          }
          
          //отправка тейка
          handlers.NeanonTxt(chatID, bot, adminsChatID, userID, msgText, userName) 
          //возврат в стартовое меню
          handlers.Start(chatID, bot)
          break
        } 
      }

    }
  }
}
  
func logTxt(text string, username string) {
  log.Printf("сообщение от @%s: %s", username, text)
}
