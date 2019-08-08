package main

import (
  "github.com/taisa831/sandbox-gin/db"
  "github.com/taisa831/sandbox-gin/router"
)

func main() {
  dbConn := db.Init()
  router.Router(dbConn)
}
