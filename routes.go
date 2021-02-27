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

var endpoints = []Endpoint{ MakeEndpoint("/user/login", 1, services.LoginUser),
                            MakeEndpoint("/user/signup", 1, services.SignUpUser),
                            //MakeEndpoint("/app/upload", 1, services.UploadApp),
                            MakeEndpoint("/app/get", 0, services.GetAllApps),
                            MakeEndpoint("/best/:category", 0, services.GetBest),
                            MakeEndpoint("/app/get/:id", 0, services.GetAppById),
                            MakeEndpoint("/app/comments/:id", 0, services.GetCommentsForApp),
                            MakeEndpoint("/app/comments", 1, services.PostComment),
                            MakeEndpoint("/app/download/:id", 0, services.DownloadApp),
                           }

func DefineRoutes(g *gin.Engine) {
  for i := range(endpoints) {
    if endpoints[i].Method == 0 {
      g.GET(endpoints[i].Path, endpoints[i].Handler)
    } else if endpoints[i].Method == 1 {
      g.POST(endpoints[i].Path, endpoints[i].Handler)
    }

    //for CORS bypass
    g.OPTIONS(endpoints[i].Path, func(c *gin.Context){
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, *")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
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
