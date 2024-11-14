package main

import (
	"log"
	"os"

	"cf-bot/internal/handlers"
	"cf-bot/internal/begin"

	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("cant load .env")
  }

  token := os.Getenv("TOKEN")
  adminsChatIDstr := os.Getenv("ADM_CHAT")

  adminsChatID, updates, bot := begin.Start(token, adminsChatIDstr)

  for u := range updates {
    //если нет сообщения пропускаем итерацию
    if u.Message == nil {
      continue
    }

    //создаем переменные
    chatID, msgText, userName := begin.CreateVars(u)

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
        handlers.AnonTxt(updates, bot, adminsChatID)
      case "неанон":
        logTxt(msgText, userName)
        handlers.TakeTxt(chatID, bot)
        handlers.NeanonTxt(updates, bot, adminsChatID)
      }
    }
  }
}
  
func logTxt(text string, username string) {
  log.Printf("сообщение от @%s: %s", username, text)
}

func logPhoto(username string) {
  log.Printf("фото от @%s", username)
}
