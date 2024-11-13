package keyboards

import (
  tg "github.com/OvyFlash/telegram-bot-api"
)

var StartKB = tg.NewReplyKeyboard(
  tg.NewKeyboardButtonRow(
    tg.NewKeyboardButton("анон"),
    ),
  tg.NewKeyboardButtonRow(
    tg.NewKeyboardButton("неанон"),
    ),
  )
var TakeKB = tg.NewReplyKeyboard(
  tg.NewKeyboardButtonRow(
    tg.NewKeyboardButton("не хочу отправлять тейк"),
    ),
  )
