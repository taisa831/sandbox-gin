package main

import (
  "github.com/taisa831/gin-sandbox/db"
  "github.com/taisa831/gin-sandbox/router"
)

func main() {
  dbConn := db.Init()
  router.Router(dbConn)
}
