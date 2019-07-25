package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  "github.com/taisa831/gin-sandbox/models"
  "net/http"
  "strconv"
)

type TodoHandler struct {
  Db *gorm.DB
}

func (h *TodoHandler) GetAll(c *gin.Context) {

  var todos []models.Todo
  h.Db.Find(&todos)

  c.HTML(http.StatusOK, "index.html", gin.H{
    "todos": todos,
  })
}

func (h *TodoHandler) CreateTask(c *gin.Context) {
  text, _ := c.GetPostForm("text")
  status, _ := c.GetPostForm("status")
  istatus, _ := strconv.ParseUint(status, 10, 32)

  h.Db.Create(&models.Todo{Text: text, Status: istatus})
  c.Redirect(http.StatusMovedPermanently, "/todo")
}

func (h *TodoHandler) EditTask(c *gin.Context) {
  todo := models.Todo{}
  id := c.Param("id")
  h.Db.First(&todo, id)
  c.HTML(http.StatusOK, "edit.html", gin.H{
    "todo": todo,
  })
}

func (h *TodoHandler) UpdateTask(c *gin.Context) {
  todo := models.Todo{}
  id := c.Param("id")
  text, _ := c.GetPostForm("text")
  status, _ := c.GetPostForm("status")
  istatus, _ := strconv.ParseUint(status, 10, 32)

  h.Db.First(&todo, id)
  todo.Text = text
  todo.Status = istatus
  h.Db.Save(&todo)
  c.Redirect(http.StatusMovedPermanently, "/todo")
}

func (h *TodoHandler) DeleteTask(c *gin.Context) {
  todo := models.Todo{}
  id := c.Param("id")
  h.Db.First(&todo, id)
  h.Db.Delete(&todo)
  c.Redirect(http.StatusMovedPermanently, "/todo")
}
