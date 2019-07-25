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

  r.GET("/go/todo", todoHandler.GetAll) // 一覧画面
  r.POST("go/todo", todoHandler.CreateTask) // 新規作成
  r.GET("/go/todo/:id", todoHandler.EditTask) // 編集画面
  r.POST("/go/todo/edit/:id", todoHandler.UpdateTask) // 更新
  r.POST("/go/todo/delete/:id", todoHandler.DeleteTask) // 削除

  apiV1 := r.Group("/api/v1")
  {
    apiTodoHandler := v1.TodoHandler{
      Db: dbConn,
    }

    apiV1.GET("/go/todo", apiTodoHandler.GetAll)
    apiV1.POST("/go/todo", apiTodoHandler.CreateTask)
    apiV1.GET("/go/todo/:id", apiTodoHandler.EditTask)
    apiV1.POST("/go/todo/edit/:id", apiTodoHandler.UpdateTask)
    apiV1.POST("/go/todo/delete/:id", apiTodoHandler.DeleteTask)
  }

  r.Run(":9000")
}
