package router

import (
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  v1 "github.com/taisa831/gin-sandbox/api/v1"
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

  apiV1 := r.Group("/api/v1")
  {
    apiTodoHandler := v1.TodoHandler{
      Db: dbConn,
    }

    apiV1.GET("/todo", apiTodoHandler.GetAll)
    apiV1.POST("/todo", apiTodoHandler.CreateTask)
    apiV1.GET("/todo/:id", apiTodoHandler.EditTask)
    apiV1.PUT("/todo/:id", apiTodoHandler.UpdateTask)
    apiV1.POST("/todo/:id", apiTodoHandler.DeleteTask)
  }

  r.Run(":9000")
}
