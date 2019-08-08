package router

import (
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  v1 "github.com/taisa831/sandbox-gin/api/v1"
  "github.com/taisa831/sandbox-gin/controllers"
  "time"
)

func Router(dbConn *gorm.DB) {

  todoHandler := controllers.TodoHandler{
    Db: dbConn,
  }

  r := gin.Default()
  r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"PUT", "PATCH", "DELETE", "POST", "GET"},
    AllowHeaders:     []string{"Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "*"
    },
    MaxAge: 12 * time.Hour,
  }))

  r.LoadHTMLGlob("templates/*")

  r.GET("/todo", todoHandler.GetAll)                 // 一覧画面
  r.POST("todo", todoHandler.CreateTask)             // 新規作成
  r.GET("/todo/:id", todoHandler.EditTask)           // 編集画面
  r.POST("/todo/edit/:id", todoHandler.UpdateTask)   // 更新
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
    apiV1.DELETE("/todo/:id", apiTodoHandler.DeleteTask)
  }

  r.Run(":9000")
}
