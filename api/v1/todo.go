package v1

import (
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  "github.com/taisa831/gin-sandbox/models"
  "net/http"
)

type TodoHandler struct {
  Db *gorm.DB
}

func (h *TodoHandler) GetAll(c *gin.Context) {

  accessControl(c)

  var todos []models.Todo
  h.Db.Find(&todos)
  c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) CreateTask(c *gin.Context) {

  accessControl(c)

  todo := models.Todo{}

  err := c.BindJSON(&todo)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  h.Db.Create(&todo)
  c.JSON(http.StatusOK, &todo)
}

func (h *TodoHandler) EditTask(c *gin.Context) {

  accessControl(c)

  todo := models.Todo{}
  id := c.Param("id")
  h.Db.First(&todo, id)
  c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) UpdateTask(c *gin.Context) {

  accessControl(c)

  todo := models.Todo{}
  id := c.Param("id")
  h.Db.First(&todo, id)

  err := c.BindJSON(&todo)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
  }

  h.Db.Save(&todo)
  c.JSON(http.StatusOK, &todo)
}

func (h *TodoHandler) DeleteTask(c *gin.Context) {

  accessControl(c)

  todo := models.Todo{}
  id := c.Param("id")
  h.Db.First(&todo, id)

  err := h.Db.First(&todo, id).Error
  if err != nil {
    c.AbortWithStatus(http.StatusNotFound)
    return
  }

  h.Db.Delete(&todo)
  c.JSON(http.StatusOK, gin.H{
    "status": "ok",
  })
}

func accessControl(c *gin.Context) {

  if c.Request.Method == "OPTIONS" {
    headers := c.Request.Header.Get("Access-Control-Request-Headers")

    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
    c.Writer.Header().Set("Access-Control-Allow-Headers", headers)

    c.Data(200, "text/plain", []byte{})
    c.Abort()
  } else {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    c.Header("Access-Control-Max-Age", "86400")
    c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  }

}