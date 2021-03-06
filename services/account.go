package services

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "io/ioutil"
  "mydroid/models"
  "encoding/json"
  "errors"
)


func SignUpUser(c *gin.Context) {
  body, err := ioutil.ReadAll(c.Request.Body)
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
  }

  var root map[string]string

  json.Unmarshal(body, &root)

  username := root["username"]
  password := root["password"]
  c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

  if CheckUsernameExists(username) {
    c.JSON(403, gin.H{
      "result": "fail",
      "message": "This username already exitsts in database",
    })
    return
  }

  println(fmt.Sprintf("INSERT INTO account(type, name, password) values(%d, '%s', '%s')", 3, username, password))
  _, err = db.Exec(fmt.Sprintf("INSERT INTO account(type, name, password) values(%d, '%s', '%s')", 3, username, password))
  if err != nil {
    panic(err)
  }
}

func CheckUserAndPassword(username string, password string) (bool, int, error) {
  un := ""
  pwd := ""
  id := 0
  err := db.QueryRow(fmt.Sprintf("SELECT id, name, password FROM account where name='%s'", username)).Scan(&id, &un, &pwd)

  if un == "" {
    return false, -1, err

  }

  if err != nil {
    return false, -1, err // proper error handling instead of panic in your app
  }

  if pwd == password {
    return true, id, nil
  }
  return false,-1, nil
}


func LoginUser(c *gin.Context) {
  body, err := ioutil.ReadAll(c.Request.Body)
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
  }


  var root map[string]string

  json.Unmarshal(body, &root)

  username := root["username"]
  password := root["password"]
  println(username, password)
  c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

  if username != "" || password != "" {
     authenticated, user_id, _ := CheckUserAndPassword(username, password)
     if authenticated {
       token := GenerateUserToken(user_id)
       c.SetCookie("X-AUTH", token, 3600, "/", "localhost", false, true)

       c.JSON(200, gin.H{
         "result": "success",
         "message": gin.H{
           "user_id": user_id,
         },
       })
     } else {
       c.JSON(403, gin.H{
         "result": "fail",
         "message": "The given credentials were not right",
       })
     }
  }

}

func GetAccountByID(pid int) (models.Account, error) {
  var acc models.Account
  query := fmt.Sprintf("SELECT id,name,type,password FROM account WHERE id=%d", pid);
  println(query)
  rows, err := db.Query(query)
  if err != nil {
    return acc, err
  }

  if !rows.Next() {
    return acc, errors.New("No accounts found with provided ID")
  }


  rows.Scan(&acc.ID, &acc.Name, &acc.Type, &acc.Password)
  return acc, nil
}

func CheckUsernameExists(username string) bool {
  query := fmt.Sprintf("SELECT * FROM account WHERE name='%s'", username)
  rows, err := db.Query(query)
  if err != nil {
    fmt.Println(err)
  }
  return rows.Next()
}
