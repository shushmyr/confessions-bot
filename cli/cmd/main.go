package main

import (
  "fmt"
  "cf-bot/internal/users"
)

func main() {
  fmt.Println("welcome to bot-cf cli")

  for {
    fmt.Println("choose what you going to do")
    var choose int
    fmt.Println("1 - ban, 2 - get users")
    fmt.Scan(&choose)

    switch choose {
    case 1:
      fmt.Println("enter id of user which you want to ban")
      var choose int 
      fmt.Scan(&choose)

      //users.Ban() 
    case 2:
      fmt.Println(users.GetAll())
    }
  }
}
