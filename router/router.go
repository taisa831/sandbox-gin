package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	v1 "github.com/taisa831/sandbox-gin/api/v1"
	"github.com/taisa831/sandbox-gin/controllers"
	"github.com/taisa831/sandbox-gin/middleware"
	"log"
	"time"

	"github.com/appleboy/gin-jwt"
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

	authMiddleware, err := middleware.GetJwtAuth()
	if err != nil {
		panic(err.Error())
	}

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

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

		signUpHandler := controllers.SignUpHandler{
			Db: dbConn,
		}

		apiV1.POST("/signup", signUpHandler.SignUp)
		apiV1.POST("/login", authMiddleware.LoginHandler)
		apiV1.GET("/todo", apiTodoHandler.GetAll)
		apiV1.POST("/todo", apiTodoHandler.CreateTask)
		apiV1.GET("/todo/:id", apiTodoHandler.EditTask)
		apiV1.PUT("/todo/:id", apiTodoHandler.UpdateTask)
		apiV1.DELETE("/todo/:id", apiTodoHandler.DeleteTask)

		apiV1.Use(authMiddleware.MiddlewareFunc())
		{
			apiV1.GET("/hello", helloHandler)
		}
	}

	if err := r.Run(":9000"); err != nil {
		panic(err)
	}
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(middleware.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*middleware.User).UserName,
		"text":     "Hello World.",
	})
}
