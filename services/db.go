package services

import (
  "database/sql"
)

var db *sql.DB

func InitializeDb() {
  datab, err := sql.Open("mysql", "root:root@tcp(localhost)/mydroid")
  if err != nil {
    panic(err)
  }
  db = datab

}
