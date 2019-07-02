package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/taisa831/gin-sandbox/models"
)

func Init() *gorm.DB {
  db, err := gorm.Open("mysql", "gorm:gorm@/sandbox_gin?charset=utf8mb4&parseTime=True&loc=Local")
  if err != nil {
    panic("データベースへの接続に失敗しました")
  }
  db.LogMode(true)
  db.AutoMigrate(&models.Todo{})

  return db
}
