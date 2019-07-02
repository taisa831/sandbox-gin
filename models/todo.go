package models

import "github.com/jinzhu/gorm"

type Todo struct {
 gorm.Model
 Text string
 Status uint64
}
