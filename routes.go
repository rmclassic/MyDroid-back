package main

import (
  "github.com/gin-gonic/gin"
  "mydroid/services"
  //"encoding/json"
  //"database/sql"
)

type Endpoint struct {
  Path string
  Method int
  Handler gin.HandlerFunc
}

<<<<<<< HEAD
var endpoints = []Endpoint{ MakeEndpoint("/user/login", 1, services.LoginUser),
                            MakeEndpoint("/user/signup", 1, services.SignUpUser),
                            MakeEndpoint("/app/upload", 1, services.UploadApp),
                           }
=======
var endpoints = []Endpoint{ MakeEndpoint("/login", 1, services.LoginUser),
                            MakeEndpoint("/signup", 1, services.SignUpUser),
                            MakeEndpoint("/app/get", 0, services.GetAllApps)}
>>>>>>> f357b18fffd2ae6d032c72162e97945af297d1d4

func DefineRoutes(g *gin.Engine) {
  for i := range(endpoints) {
    if endpoints[i].Method == 0 {
      g.GET(endpoints[i].Path, endpoints[i].Handler)
    } else if endpoints[i].Method == 1 {
      g.POST(endpoints[i].Path, endpoints[i].Handler)
    }

    //for CORS bypass
    g.OPTIONS(endpoints[i].Path, func(c *gin.Context){
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

        c.String(200, "pong")
      })
  }
}





func MakeEndpoint(path string, method int, callback func(c *gin.Context)) Endpoint {
  var f Endpoint
  f.Path = path
  f.Method = method
  f.Handler = callback
  return f
}
