package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"cf-bot/internal/handlers"
	"cf-bot/internal/texts"
	"cf-bot/internal/users"

	tg "github.com/OvyFlash/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")
  if err != nil {
    panic("can't load token")
  }

  token := os.Getenv("TOKEN")
  adminsChatID := int64(-1002269839756)

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
    userID := u.SentFrom().ID

    banned := users.CheckIfBanned(userID)
    if banned == true {
      msg := tg.NewMessage(chatID, "ты забанен")
      bot.Send(msg)
      logTxt(msgText, userName)
      log.Printf("забаненный пользователь @%s писал сообщение", userName)
      break
    }


    //обрабатываем текст
    if msgText != "" {
      //если от админов
      if chatID == adminsChatID {
        switch msgText {
        case "/response":
          logTxt(msgText, userName)

          admID := u.SentFrom().ID
          msg := tg.NewMessage(chatID, texts.AdmResponse)
          bot.Send(msg)

          for u := range updates {
            if u.SentFrom().ID != admID {
              continue
            }
            
            msgText := strings.Split(u.Message.Text, ",")
            userIDi, err := strconv.Atoi(msgText[0])
            if err != nil {
              log.Println("не удалось сконвертировать userID в число")
            }

            userID := int64(userIDi)
            response := msgText[1]
            adminUserName := u.SentFrom().UserName

            handlers.AdmResponse(userID, adminsChatID, bot, response, adminUserName)
            break
          }
        }
      }
      switch msgText {
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
