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
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

  println(fmt.Sprintf("INSERT INTO account(type, name, password) values(%d, '%s', '%s')", 3, username, password))
  _, err = db.Exec(fmt.Sprintf("INSERT INTO account(type, name, password) values(%d, '%s', '%s')", 3, username, password))
  if err != nil {
    panic(err)
  }
}

func CheckUserAndPassword(username string, password string) (bool, error) {
  un := ""
  pwd := ""
  err := db.QueryRow(fmt.Sprintf("SELECT name, password FROM account where name='%s'", username)).Scan(&un, &pwd)

  if un == "" {
    return false, err

  }

  if err != nil {
    return false, err // proper error handling instead of panic in your app
  }

  if pwd == password {
    return true, nil
  }
  return false, nil
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
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

  if username != "" || password != "" {
     authenticated, err := CheckUserAndPassword(username, password)
     if authenticated {
       c.JSON(200, gin.H{
         "result": "success",
         "message": "",
       })
     } else {
       c.JSON(200, gin.H{
         "result": "fail",
         "message": err,
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
