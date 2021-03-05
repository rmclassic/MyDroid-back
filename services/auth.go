package services

import (
  "fmt"
  "math/rand"
  "time"
)

func ValidateUserToken(userid int, token string) bool {
  query := fmt.Sprintf("SELECT token FROM token WHERE uid=%d", userid)
  rows, _ := db.Query(query)
  if !rows.Next() {
    return false
  }

  var utoken string
  rows.Scan(&utoken)

  if utoken == token {
    return true
  } else {
    return false
  }
}

func GenerateUserToken(userid int) string {
  rand.Seed(time.Now().UnixNano())
  var letterRunes = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

  b := make([]rune, 19)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }

    query := ""
    if GetTokenOfUser(userid) == "" {
      query = fmt.Sprintf("INSERT INTO token values (%d, '%s')", userid, string(b))
    } else {
      query = fmt.Sprintf("UPDATE token SET token='%s' WHERE uid=%d", string(b), userid)
    }

    _, err := db.Exec(query)
    fmt.Println(err)
    fmt.Println(query)
    return string(b)
}

func GetUserByToken(token string) int {
  query := fmt.Sprintf("SELECT uid FROM token WHERE token='%s'", token)
  rows, err := db.Query(query)
  if err != nil {
    fmt.Println(err)
    return -1
  }

  if !rows.Next() {
    return -1
  }

  var userid int
  rows.Scan(&userid)
  return userid
}

func GetTokenOfUser(userid int) string {
  query := fmt.Sprintf("SELECT token FROM token WHERE uid=%d", userid)
  rows, err := db.Query(query)
  if err != nil {
    return ""
  }

  if !rows.Next() {
    return ""
  }

  token := ""
  rows.Scan(&token)
  return token
}
