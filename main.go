package main

import (
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
)

var db *sql.DB

func main() {
  g := gin.Default()
  datab, err := sql.Open("mysql", "root:toor@tcp(localhost)/mydroid")
  if err != nil {
    panic(err)
  }
  db = datab
  DefineRoutes(g)

  g.Run()
}
