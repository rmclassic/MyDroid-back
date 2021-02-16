package main

import (
  "github.com/gin-gonic/gin"
  "fmt"
  //"encoding/json"
  //"database/sql"
)

type Endpoint struct {
  Path string
  Method int
  Handler gin.HandlerFunc
}

var endpoints = []Endpoint{ MakeEndpoint("/login", 1, LoginUser), MakeEndpoint("/signup", 1, SignUpUser) }

func DefineRoutes(g *gin.Engine) {
  for i := range(endpoints) {
    if endpoints[i].Method == 0 {
      g.GET(endpoints[i].Path, endpoints[i].Handler)
    } else if endpoints[i].Method == 1 {
      g.POST(endpoints[i].Path, endpoints[i].Handler)
    }
  }
}

func LoginUser(c *gin.Context) {
  username := c.DefaultPostForm("username", "")
  password := c.DefaultPostForm("password", "")
  println(username, password)
  if username != "" || password != "" {
     if CheckUserAndPassword(username, password) {
       c.JSON(200, gin.H{
         "result": "success",
         "message": "",
       })
     } else {
       c.JSON(200, gin.H{
         "result": "fail",
         "message": "invalid credentials",
       })
     }
  }
}

func SignUpUser(c *gin.Context) {
  username := c.DefaultPostForm("username", "")
  password := c.DefaultPostForm("password", "")
  println(fmt.Sprintf("INSERT INTO account(type, name, password) values(%d, '%s', '%s')", 3, username, password))
  _, err := db.Exec(fmt.Sprintf("INSERT INTO account(type, name, password) values(%d, '%s', '%s')", 3, username, password))
  if err != nil {
    panic(err)
  }
}


func MakeEndpoint(path string, method int, callback func(c *gin.Context)) Endpoint {
  var f Endpoint
  f.Path = path
  f.Method = method
  f.Handler = callback
  return f
}

func CheckUserAndPassword(username string, password string) bool {
  un := ""
  pwd := ""
  err := db.QueryRow(fmt.Sprintf("SELECT name, password FROM account where name='%s'", username)).Scan(&un, &pwd)
  if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
  }

  if pwd == password {
    return true
  }
  return false
}
