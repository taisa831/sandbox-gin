package models

import (
	"time"
)

type User struct {
	ID                uint       `json:"-" gorm:"primary_key"`
	CreatedAt         time.Time  `json:"createdAt" gorm:"index"`
	UpdatedAt         time.Time  `json:"updatedAt" gorm:"index"`
	DeletedAt         *time.Time `json:"deletedAt" gorm:"index"`
	Uuid              string     `json:"uuid" gorm:"unique_index"`
	Email             string     `json:"email" gorm:"unique_index"`
	UserName          string     `json:"userName" gorm:"index"`
	Password          string     `json:"password"`
	ConfirmationToken string     `json:"confirmationToken" gorm:"index"`
}

