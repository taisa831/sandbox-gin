package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	v1 "github.com/taisa831/sandbox-gin/api/v1"
	"github.com/taisa831/sandbox-gin/controllers"
	"log"
	"time"

	"github.com/appleboy/gin-jwt"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

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

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
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

	r.Run(":9000")
}
