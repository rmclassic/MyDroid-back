package main

import (
  "fmt"
  "reflect"
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
)

var db *sql.DB

func main() {
  g := gin.Default()
  db, _ = sql.Open("mysql", "root:root@tcp(localhost)/mydroid")
  println("db isisisis", db)
  //if err != nil {
    //panic(err)
  //}
  fmt.Println(reflect.TypeOf(db))
  DefineRoutes(g)

  g.Run()
}
