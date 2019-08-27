package controllers

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/taisa831/sandbox-gin/models"
	"math/big"
	"net/http"
)

type SignUpHandler struct {
	Db *gorm.DB
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func (s *SignUpHandler) SignUp(c *gin.Context) {

	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	user.ConfirmationToken = ""
	user.Uuid = "user-" + Phase1(100)

	s.Db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"confirmation_token": &user.ConfirmationToken})
}

func Phase1(n int) string {
	buf := make([]byte, n)
	max := new(big.Int)
	max.SetInt64(int64(len(letterBytes)))
	for i := range buf {
		r, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		buf[i] = letterBytes[r.Int64()]
	}
	return string(buf)
}
