package begin

import (
	"log"
	"strconv"

	tg "github.com/OvyFlash/telegram-bot-api"
	"github.com/joho/godotenv"
)

func Start(token, adminsChatIDstr string) (int64, tg.UpdatesChannel, *tg.BotAPI) {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("cant load .env")
  }
 
  adminsChatIDi, err := strconv.Atoi(adminsChatIDstr)
  if err != nil {
    log.Fatal("can't get admins chat id")
  }
  adminsChatID := int64(adminsChatIDi)

  bot, err := tg.NewBotAPI(token)
  if err != nil {
    log.Fatal("cant create bot")
  }

  log.Printf("авторизован под %s", bot.Self.UserName)

  updateConfig := tg.NewUpdate(0)
  updateConfig.Timeout = 60

  updates := bot.GetUpdatesChan(updateConfig)
  
  return adminsChatID, updates, bot
}

func CreateVars(u tg.Update) (chatID int64, msgText, UserName string) {
  chatID = u.Message.Chat.ID
  msgText = u.Message.Text
  UserName = u.SentFrom().UserName
  
  return chatID, msgText, UserName
}
