package router

import (
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  "github.com/taisa831/gin-sandbox/controllers"
)

func Router(dbConn *gorm.DB) {

  todoHandler := controllers.TodoHandler{
    Db: dbConn,
  }

  r := gin.Default()
  r.LoadHTMLGlob("templates/*")

  r.GET("/todo", todoHandler.GetAll) // 一覧画面
  r.POST("/todo", todoHandler.CreateTask) // 新規作成
  r.GET("/todo/:id", todoHandler.EditTask) // 編集画面
  r.POST("/todo/edit/:id", todoHandler.UpdateTask) // 更新
  r.POST("/todo/delete/:id", todoHandler.DeleteTask) // 削除

  r.Run(":9000")
}
