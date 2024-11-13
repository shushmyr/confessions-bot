package users

var Users = make(map[int64]int64)
var Banned = []int64{}

func Add(userID, chatID int64) {
  x := CheckIfExist(userID, chatID)
  if x == true {
    return
  } else {
    Users[userID] = chatID
  }
}

func CheckIfExist(userID, chatID int64) bool {
  response := false
  for i := int64(0); i < int64(len(Users)); i++ {
    if Users[i] == chatID {
      response = true
      break
    }
  }

  return response
}

func GetAll() map[int64]int64 {
  return Users
}

func GetUser(userID int64) int64 {
  return Users[userID]
}

func Ban(userID int64) {
  x := CheckIfBanned(userID)
  if x == true {
    return
  } else {
    Banned = append(Banned, userID)
  } 
}

func CheckIfBanned(userID int64) bool {
  response := false
  for i := int64(0); i < int64(len(Users)); i++ {
    if Banned[i] == userID {
      response = true
      break
    }
  }

  return response
}
