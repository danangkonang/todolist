package service

import (
	"gorm.io/gorm"
)

type connection struct {
	Mysql *gorm.DB
}
