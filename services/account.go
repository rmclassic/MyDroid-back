package services

import (
  "github.com/gin-gonic/gin"
  "fmt"
)



func SignUpUser(c *gin.Context) {
  username := c.DefaultPostForm("username", "")
  password := c.DefaultPostForm("password", "")
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

  println(fmt.Sprintf("INSERT INTO account(type, name, password) values(%d, '%s', '%s')", 3, username, password))
  _, err := db.Exec(fmt.Sprintf("INSERT INTO account(type, name, password) values(%d, '%s', '%s')", 3, username, password))
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
  username := c.DefaultPostForm("username", "")
  password := c.DefaultPostForm("password", "")
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

  println(username, password)
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
