package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/taisa831/sandbox-gin/config"
  "github.com/taisa831/sandbox-gin/models"
  "os"
)

const confDir = "./config/env/"

func Init() *gorm.DB {

  appMode := os.Getenv("APP_MODE")
  if appMode == "" {
    panic("failed to get application mode, check whether APP_MODE is set.")
  }

  conf, err := config.NewConfig(confDir, appMode)
  if err != nil {
    panic(err.Error())
  }

  db, err := gorm.Open("mysql", conf.DB.User + ":" + conf.DB.Password + "@/" + conf.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local")
  if err != nil {
    panic("データベースへの接続に失敗しました")
  }
  db.LogMode(true)
  db.AutoMigrate(&models.Todo{}, &models.Todo{})

  return db
}
