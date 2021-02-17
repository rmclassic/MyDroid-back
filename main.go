package main

import (
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
  "mydroid/services"
)



func main() {
  g := gin.Default()
  services.InitializeDb()
  DefineRoutes(g)

  g.Run()
}
